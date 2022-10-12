package controller

import (
	"encoding/json"
	"net/http"
	"web-service/constant"
	"web-service/pkg/auth"
	"web-service/pkg/log"
	"web-service/pkg/router"
	usermodel "web-service/service/account/model"
	"web-service/service/channel/model"
)

type CreateNewChannel struct {
	Name string `json:"name"`
}

func CreateChannel(w http.ResponseWriter, r *http.Request) {

	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "CreateChannel: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateChannel: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	user := &usermodel.User{
		Id: payload.UserId,
	}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err = user.GetUserById()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "CreateChannel: Get users Database", err)
		return
	}

	NewChannel := &CreateNewChannel{}
	_ = json.NewDecoder(r.Body).Decode(NewChannel)

	var channel = &model.Channel{
		Name:       NewChannel.Name,
		MemberId:   user.Id,
		MemberName: user.Name,
		Role:       constant.CHANNEL_ROLE_HOST,
	}

	err = channel.CreateChannel()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateChannel", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "B.ACC.201.C1", "Create channel successfully !!!")
}
