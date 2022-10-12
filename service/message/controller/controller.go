package controller

import (
	"encoding/json"
	"net/http"
	"web-service/pkg/log"
	"web-service/pkg/router"
	"web-service/service/message/model"
)

type NewMessage struct {
	Content string `json:"content"`
}

func InsertMessage(w http.ResponseWriter, r *http.Request) {
	NewMessage := &NewMessage{}
	_ = json.NewDecoder(r.Body).Decode(NewMessage)

	var message = &model.Message{
		Content: NewMessage.Content,
	}

	var err = message.InsertMessage()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateMessage: InsertMessage", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.ACC.201.C1", "Create message successfully !!!")
}
