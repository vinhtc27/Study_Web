package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"web-service/constant"
	"web-service/pkg/auth"
	"web-service/pkg/log"
	"web-service/pkg/router"
	"web-service/pkg/utils"
	"web-service/service/channel/model"

	userModel "web-service/service/account/model"

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

	channelImageIndex := rand.Intn(len(constant.DEFAULT_CHANNEL_AVATAR_LIST))

	var channel = &model.Channel{
		Name:   newChannel.Name,
		Avatar: constant.DEFAULT_CHANNEL_AVATAR_LIST[channelImageIndex],
		Members: []model.Member{
			{
				UserId: payload.UserId,
				Role:   constant.CHANNEL_ROLE_HOST,
			},
		},
		CreatedDate: utils.Timestamp(),
		UpdatedDate: utils.Timestamp(),
	}
	err = channel.CreateChannel()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateChannel", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccessWithData(w, "B.CHA.201.C1", "Create channel successfully !!!", channel)
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
			message := &model.Message{
				Type:      model.DeleteType,
				ChannelId: channel.Id,
				SenderId:  hostId,
				Content:   strconv.Itoa(member.UserId),
				Timestamp: utils.Timestamp(),
			}
			Broadcast <- message
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

func GetChannelById(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "GetChannelById: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "GetChannelById: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "GetChannelById: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}
	err = channel.GetChannelById()
	if err != nil {
		log.Println(log.LogLevelDebug, "GetChannelById: GetChannelById", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	for _, member := range channel.Members {
		if payload.UserId == member.UserId {
			router.ResponseSuccessWithData(w, "B.CHA.200.C2", "Get channel by id successfully", channel)
			return
		}
	}
	router.ResponseBadRequest(w, "B.CHA.200.C9", "You are not channel's member")
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

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateChannelById", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		newChannel := &CreateNewChannel{}
		err = json.NewDecoder(r.Body).Decode(newChannel)
		if err != nil {
			log.Println(log.LogLevelDebug, "Name null", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		channel.Name = newChannel.Name
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

	newMemberEmail := r.URL.Query().Get("email")
	user := &userModel.User{
		Email: newMemberEmail,
	}

	err = user.GetUserByEmail()
	if err != nil {
		log.Println(log.LogLevelDebug, "AddChannelMember: GetUserByEmail", err)
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
		if hostId != user.Id {
			err = channel.AddNewMember(&model.Member{UserId: user.Id, Role: constant.CHANNEL_ROLE_MEMBER})
			if err != nil {
				log.Println(log.LogLevelDebug, "AddChannelMember: AddNewMember", err)
				router.ResponseInternalError(w, err.Error())
				return
			}
			message := &model.Message{
				Type:      model.AddType,
				ChannelId: channel.Id,
				SenderId:  hostId,
				Content:   strconv.Itoa(user.Id),
				Timestamp: utils.Timestamp(),
			}
			Broadcast <- message
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
			message := &model.Message{
				Type:      model.DeleteType,
				ChannelId: channel.Id,
				SenderId:  hostId,
				Content:   strconv.Itoa(deleteUserId),
				Timestamp: utils.Timestamp(),
			}
			Broadcast <- message
			router.ResponseSuccess(w, "B.CHA.200.C7", "Delete member from channel successfully")
			return
		} else {
			router.ResponseSuccess(w, "B.CHA.200.C8", "Host cannot be deleted")
			return
		}
	} else if hostId != payload.UserId && deleteUserId == payload.UserId {
		err = channel.DeleteMember(&model.Member{UserId: deleteUserId, Role: constant.CHANNEL_ROLE_MEMBER})
		if err != nil {
			log.Println(log.LogLevelDebug, "DeleteChannelMember: DeleteMember", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		message := &model.Message{
			Type:      model.DeleteType,
			ChannelId: channel.Id,
			SenderId:  payload.UserId,
			Content:   strconv.Itoa(deleteUserId),
			Timestamp: utils.Timestamp(),
		}
		Broadcast <- message
		router.ResponseSuccess(w, "B.CHA.200.C7", "Leave channel successfully")
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}

type CreateNewTaskColumn struct {
	Title            string `json:"title"`
	TaskColumnDetail any    `json:"taskColumnDetail"`
}

type RemoveTaskColumn struct {
	Title string `json:"title"`
}

func AddTaskColumn(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "AddTaskColumn: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "AddTaskColumn: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "AddTaskColumn: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	newTaskColumn := &CreateNewTaskColumn{}
	err = json.NewDecoder(r.Body).Decode(newTaskColumn)
	if err != nil {
		log.Println(log.LogLevelDebug, "AddTaskColumn: Decode(newTaskColumn)", err.Error())
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}

	var taskColumn = &model.TaskColumn{
		Title:            newTaskColumn.Title,
		TaskColumnDetail: newTaskColumn.TaskColumnDetail,
	}

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "AddChannelMember: GetChannelHostId", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		err = channel.AddTaskColumn(taskColumn)
		if err != nil {
			log.Println(log.LogLevelDebug, "AddTaskColumn: AddTaskColumn", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		router.ResponseSuccess(w, "B.CHA.200.C5", "Add task column to channel successfully")
		return
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}

func DeleteTaskColumn(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "DeleteTaskColumn: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteTaskColumn: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteTaskColumn: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	deleteTaskColumn := &RemoveTaskColumn{}
	err = json.NewDecoder(r.Body).Decode(deleteTaskColumn)
	if err != nil {
		log.Println(log.LogLevelDebug, "Task Column null", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "DeleteTaskColumn: GetChannelHostId", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		err = channel.DeleteTaskColumnByTitle(deleteTaskColumn.Title)
		if err != nil {
			log.Println(log.LogLevelDebug, "DeleteTaskColumn: DeleteTaskColumn", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		router.ResponseSuccess(w, "B.CHA.200.C7", "Delete task column from channel successfully")
		return
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}

func UpdateTaskColumn(w http.ResponseWriter, r *http.Request) {
	//Get Parameters JWT claims header
	claims, err := auth.GetJWTClaims(r.Header.Get("X-JWT-Claims"))
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "UpdateTaskColumn: GetJWTClaims", err)
		return
	}

	// Get payload data from claims
	var payload = &auth.Payload{}
	err = payload.GetDataFromClaims(claims)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateTaskColumn: GetDataFromClaims", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	channelId, err := strconv.Atoi(chi.URLParam(r, "channelId"))
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateTaskColumn: strconv.Atoi(chi.URLParam(r, \"channelId\"))", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	newTaskColumn := &CreateNewTaskColumn{}
	err = json.NewDecoder(r.Body).Decode(newTaskColumn)
	if err != nil {
		log.Println(log.LogLevelDebug, "Name null", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	var channel = &model.Channel{
		Id: channelId,
	}

	newTaskColumnDetail, err := json.Marshal(newTaskColumn.TaskColumnDetail)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateTaskColumn: json.Marshal(newTaskColumn.TaskColumnDetail)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	var taskColumn = &model.TaskColumn{
		Title:            newTaskColumn.Title,
		TaskColumnDetail: newTaskColumnDetail,
	}

	hostId, err := channel.GetChannelHostId()
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateTaskColumn: GetChannelHostId", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	if hostId == payload.UserId {
		err = channel.UpdateTaskColumn(taskColumn)
		if err != nil {
			log.Println(log.LogLevelDebug, "UpdateTaskColumn: UpdateTaskColumn", err)
			router.ResponseInternalError(w, err.Error())
			return
		}
		router.ResponseSuccess(w, "B.CHA.200.C5", "Update task column to channel successfully")
		return
	} else {
		router.ResponseBadRequest(w, "B.CHA.400.C3", "You are not channel's host")
		return
	}
}
