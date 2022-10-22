package controller

import (
	"encoding/json"
	"net/http"
	"web-service/pkg/log"
	"web-service/pkg/router"
	usermodel "web-service/service/account/model"
	"web-service/service/message/model"
)

type NewMessage struct {
	SenderId  int    `json:"senderId"`
	Content   string `json:"content"`
	ChannelId int    `json:"channelId"`
}

type ModifiedMessage struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type DeletedMessage struct {
	Id        int `json:"id"`
	ChannelId int `json:"channelId"`
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	NewMessage := &NewMessage{}
	_ = json.NewDecoder(r.Body).Decode(NewMessage)

	var user = &usermodel.User{
		Id: NewMessage.SenderId,
	}

	var err = user.GetUserById()
	var message = &model.Message{
		SenderId:   user.Id,
		SenderName: user.Name,
		Content:    NewMessage.Content,
		ChannelId:  NewMessage.ChannelId,
	}

	err = message.CreateMessage()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateMessage", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.ACC.201.C1", "Create message successfully !!!")
}

func ModifyMessage(w http.ResponseWriter, r *http.Request) {
	ModifiedMessage := &ModifiedMessage{}
	_ = json.NewDecoder(r.Body).Decode(ModifiedMessage)

	var mess = &model.Message{
		Id:      ModifiedMessage.Id,
		Content: ModifiedMessage.Content,
	}

	var err = mess.ModifyMessage()
	if err != nil {
		log.Println(log.LogLevelDebug, "ModifyMessage", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.ACC.201.C1", "Modify message successfully !!!")
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	DeletedMessage := &DeletedMessage{}
	_ = json.NewDecoder(r.Body).Decode(DeletedMessage)

	var mess = &model.Message{
		Id:        DeletedMessage.Id,
		ChannelId: DeletedMessage.ChannelId,
	}

	var err = mess.DeleteMessage()
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteMessage", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.ACC.201.C1", "Delete message successfully !!!")
}
