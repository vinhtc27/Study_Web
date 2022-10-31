package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"web-service/pkg/auth"
	"web-service/pkg/router"
	"web-service/pkg/utils"
	"web-service/service/channel/model"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]int)
	broadcast = make(chan *model.Message)

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func init() {
	go BroadcastMessages()
}

func HandlerChannelWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
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
		fmt.Println("Web Socket connection established")

		clients[client] = channelId

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
				broadcast <- message
				err = channel.UpdateNewMessage(message)
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
		message := <-broadcast
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
