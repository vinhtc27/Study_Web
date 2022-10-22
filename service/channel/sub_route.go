package channel

import (
	"github.com/go-chi/chi"
	"web-service/pkg/auth"
	"web-service/service/channel/controller"
)

var ChannelSubRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {
	ChannelSubRoute.Group(func(_ chi.Router) {
		ChannelSubRoute.With(auth.JWT).Post("/createChannel", controller.CreateChannel)
		ChannelSubRoute.With(auth.JWT).Delete("/deleteChannel", controller.DeleteChannel)
		ChannelSubRoute.With(auth.JWT).Post("/addNewMember", controller.AddNewMemberToChannel)
		ChannelSubRoute.With(auth.JWT).Patch("/changeChannelName", controller.ChangeNameChannel)
	})
}
