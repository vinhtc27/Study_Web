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

type AddNewMember struct {
	Name       string `json:"name"`
	MemberId   int    `json:"memberId"`
	MemberName string `json:"memberName"`
}

type DeleteAChannel struct {
	Name string `json:"name"`
}

type ChangeChannelName struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
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
	err = user.GetUserById()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "CreateChannel: Get users Database", err)
		return
	}

	newChannel := &CreateNewChannel{}
	err = json.NewDecoder(r.Body).Decode(newChannel)
	if err != nil {
		log.Println(log.LogLevelDebug, "Name null", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Name:       newChannel.Name,
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

func DeleteChannel(w http.ResponseWriter, r *http.Request) {

	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "DeleteChannel: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannel: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	user := &usermodel.User{
		Id: payload.UserId,
	}
	err = user.GetUserById()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "DeleteChannel: Get users Database", err)
		return
	}

	aChannel := &DeleteAChannel{}
	err = json.NewDecoder(r.Body).Decode(aChannel)
	if err != nil {
		log.Println(log.LogLevelDebug, "Name null", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Name:       aChannel.Name,
		MemberName: user.Name,
		Role:       constant.CHANNEL_ROLE_HOST,
	}
	err = channel.DeleteChannel()

	if err != nil {
		log.Println(log.LogLevelDebug, "Delete Channel", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "B.ACC.201.C1", "Delete channel successfully !!!")
}

func AddNewMemberToChannel(w http.ResponseWriter, r *http.Request) {

	newMember := &AddNewMember{}
	_ = json.NewDecoder(r.Body).Decode(newMember)

	user := &usermodel.User{
		Id: newMember.MemberId,
	}
	err := user.GetUserById()

	if err != nil {
		log.Println(log.LogLevelDebug, "Get user by id", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Name:       newMember.Name,
		MemberId:   user.Id,
		MemberName: user.Name,
		Role:       constant.CHANNEL_ROLE_MEMBER,
	}

	err = channel.AddMemberToChannel()

	if err != nil {
		log.Println(log.LogLevelDebug, "Add new member to channel", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "B.ACC.201.C1", "Add new member to channel successfully !!!")
}

func ChangeNameChannel(w http.ResponseWriter, r *http.Request) {

	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ChangeNameChannel: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "ChangeNameChannel: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	user := &usermodel.User{
		Id: payload.UserId,
	}
	err = user.GetUserById()
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "ChangeNameChannel: Get users Database", err)
		return
	}

	aChannel := &ChangeChannelName{}
	err = json.NewDecoder(r.Body).Decode(aChannel)
	if err != nil {
		log.Println(log.LogLevelDebug, "Name null", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Name:       aChannel.Name,
		Id:         aChannel.Id,
		MemberName: user.Name,
		Role:       constant.CHANNEL_ROLE_HOST,
	}
	err = channel.ChangeNameChannel()

	if err != nil {
		log.Println(log.LogLevelDebug, "Change Name Channel", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	router.ResponseSuccess(w, "B.ACC.201.C1", "Change name channel successfully !!!")
}
