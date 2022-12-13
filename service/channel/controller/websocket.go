package controller

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
	"web-service/pkg/auth"
	"web-service/pkg/crypt"
	"web-service/pkg/router"
	"web-service/pkg/utils"
	"web-service/service/channel/model"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]int)
	Broadcast = make(chan *model.Message)

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func init() {
	go BroadcastMessages()
	go Ping()
}

func HandlerChannelWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	authClaims, err := auth.JwtClaims(r.URL.Query().Get("token"))
	if err != nil {
		router.ResponseBadRequest(w, "", err.Error())
		return
	}

	claimsEncrypted, err := crypt.EncryptWithRSA(authClaims["data"].(string))
	if err != nil {
		router.ResponseBadRequest(w, "", err.Error())
		return
	}
	claims, err := auth.GetJWTClaims(claimsEncrypted)
	if err != nil {
		router.ResponseBadRequest(w, "", err.Error())
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(r.URL.Query().Get("channelId"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	channel := &model.Channel{
		Id: channelId,
	}
	err = channel.GetChannelById()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	isMember := false
	for _, member := range channel.Members {
		if member.UserId == payload.UserId {
			isMember = true
		}
	}

	if !isMember {
		router.ResponseBadRequest(w, "B.CHA.WS.C1", "You must be a channel's member")
		return
	} else {
		client, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			router.ResponseInternalError(w, err.Error())
			return
		}
		defer client.Close()

		clients[client] = channelId
		sort.Slice(channel.Messages, func(i, j int) bool {
			return channel.Messages[i].Timestamp < channel.Messages[j].Timestamp
		})

		for _, oldMessage := range channel.Messages {
			err = client.WriteJSON(oldMessage)
			if err != nil {
				client.Close()
				delete(clients, client)
				break
			}
		}

		for {
			message := &model.Message{
				Type:      model.ChatType,
				SenderId:  payload.UserId,
				ChannelId: channelId,
				Timestamp: utils.Timestamp(),
			}

			err := client.ReadJSON(&message)
			if err != nil {
				router.ResponseInternalError(w, err.Error())
				delete(clients, client)
				break
			}

			go func(message *model.Message) {
				Broadcast <- message
				err = channel.AddMessage(message)
				if err != nil {
					router.ResponseInternalError(w, err.Error())
					delete(clients, client)
					return
				}
			}(message)
		}
	}

}

func BroadcastMessages() {
	for {
		message := <-Broadcast
		for client, channelId := range clients {
			if message.ChannelId == channelId {
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func Ping() {
	for {
		for client, channelId := range clients {
			event := &model.Message{
				Type:      model.PingType,
				ChannelId: channelId,
			}
			err := client.WriteJSON(event)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		time.Sleep(time.Second * 30)
	}
}
