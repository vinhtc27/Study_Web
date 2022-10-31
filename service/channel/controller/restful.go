package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-service/constant"
	"web-service/pkg/auth"
	"web-service/pkg/log"
	"web-service/pkg/router"
	"web-service/service/channel/model"

	"github.com/go-chi/chi"
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

	newChannel := &CreateNewChannel{}
	err = json.NewDecoder(r.Body).Decode(newChannel)
	if err != nil {
		log.Println(log.LogLevelDebug, "Name null", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Name:   newChannel.Name,
		Avatar: constant.DEFAULT_CHANNE_AVATAR,
		Members: []model.Member{
			{
				UserId: payload.UserId,
				Role:   constant.CHANNEL_ROLE_HOST,
			},
		},
	}
	err = channel.CreateChannel()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateChannel", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.CHA.201.C1", "Create channel successfully !!!")
}

func DeleteChannelById(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "DeleteChannelById: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelById: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelById: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelById: GetChannelHostId", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		err := channel.GetChannelById()
		if err != nil {
			log.Println(log.LogLevelDebug, "DeleteChannelById: GetChannelById", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		for _, member := range channel.Members {
			err = channel.DeleteMember(&member)
			if err != nil {
				log.Println(log.LogLevelDebug, "DeleteChannelById: DeleteMember", err)
				router.ResponseInternalError(w, err.Error())
				return
			}
		}
		err = channel.DeleteChannelById()
		if err != nil {
			log.Println(log.LogLevelDebug, "DeleteChannelById: DeleteChannelById", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		router.ResponseSuccess(w, "B.CHA.200.C2", "Delete channel successfully")
		return
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}

func UpdateChannelById(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "UpdateChannelById: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateChannelById: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateChannelById: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}
	_ = json.NewDecoder(r.Body).Decode(&channel)

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateChannelById", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		err = channel.UpdateChannelById()
		if err != nil {
			log.Println(log.LogLevelDebug, "UpdateChannelById: UpdateChannelById", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		router.ResponseSuccess(w, "B.CHA.200.C4", "Update channel successfully")
		return
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}

func AddChannelMember(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "AddChannelMember: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "AddChannelMember: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "AddChannelMember: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	newUserId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "AddChannelMember: strconv.Atoi(r.URL.Query().Get(\"userId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}
	_ = json.NewDecoder(r.Body).Decode(&channel)

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "AddChannelMember: GetChannelHostId", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		if hostId != newUserId {
			err = channel.AddNewMember(&model.Member{UserId: newUserId, Role: constant.CHANNEL_ROLE_MEMBER})
			if err != nil {
				log.Println(log.LogLevelDebug, "AddChannelMember: AddNewMember", err)
				router.ResponseInternalError(w, err.Error())
				return
			}
			router.ResponseSuccess(w, "B.CHA.200.C5", "Add member to channel successfully")
			return
		} else {
			router.ResponseSuccess(w, "B.CHA.400.C6", "Host cannot be added to channel")
			return
		}
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}

func DeleteChannelMember(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "DeleteChannelMember: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelMember: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelMember: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	deleteUserId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelMember: strconv.Atoi(r.URL.Query().Get(\"userId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}
	_ = json.NewDecoder(r.Body).Decode(&channel)

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteChannelMember: GetChannelHostId", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		if hostId != deleteUserId {
			err = channel.DeleteMember(&model.Member{UserId: deleteUserId, Role: constant.CHANNEL_ROLE_MEMBER})
			if err != nil {
				log.Println(log.LogLevelDebug, "DeleteChannelMember: DeleteMember", err)
				router.ResponseInternalError(w, err.Error())
				return
			}
			router.ResponseSuccess(w, "B.CHA.200.C7", "Delete member from channel successfully")
			return
		} else {
			router.ResponseSuccess(w, "B.CHA.200.C8", "Host cannot be deleted")
			return
		}
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}
