package channel

import (
	"net/http"
	"web-service/pkg/auth"
	"web-service/service/channel/controller"

	"github.com/go-chi/chi"
)

var ChannelSubRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {
	ChannelSubRoute.Group(func(_ chi.Router) {
		ChannelSubRoute.With(auth.JWT).Post("/create-channel", controller.CreateChannel)
		ChannelSubRoute.With(auth.JWT).Delete("/channelId={channelId}", controller.DeleteChannelById)
		ChannelSubRoute.With(auth.JWT).Patch("/channelId={channelId}", controller.UpdateChannelById)
		ChannelSubRoute.With(auth.JWT).Patch("/add/channelId={channelId}", controller.AddChannelMember)
		ChannelSubRoute.With(auth.JWT).Post("/delete/channelId={channelId}", controller.DeleteChannelMember)
		ChannelSubRoute.With(auth.JWT).Patch("/addTaskColumn/channelId={channelId}", controller.AddTaskColumn)
		ChannelSubRoute.With(auth.JWT).Post("/deleteTaskColumn/channelId={channelId}", controller.DeleteTaskColumn)
		ChannelSubRoute.With(auth.JWT).Handle("/chat/channelId={channelId}", http.HandlerFunc(controller.HandlerChannelWebSocket))
	})
}
