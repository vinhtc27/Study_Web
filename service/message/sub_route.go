package message

import (
	"github.com/go-chi/chi"
	"web-service/service/message/controller"
)

var MessageSubRoute = chi.NewRouter()

// Init package with sub-router for message service
func init() {
	MessageSubRoute.Group(func(_ chi.Router) {
		//MessageSubRoute.With(auth.JWT).Patch("/message", controller.UpdateCurrentProfile)
		MessageSubRoute.Post("/createMessage", controller.CreateMessage)
		MessageSubRoute.Patch("/modifyMessage", controller.ModifyMessage)
		MessageSubRoute.Delete("/deleteMessage", controller.DeleteMessage)
	})
}
