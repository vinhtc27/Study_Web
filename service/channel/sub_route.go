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
		ChannelSubRoute.With(auth.JWT).Get("/channelId={channelId}", controller.GetChannelById)
		ChannelSubRoute.With(auth.JWT).Delete("/channelId={channelId}", controller.DeleteChannelById)
		ChannelSubRoute.With(auth.JWT).Patch("/channelId={channelId}", controller.UpdateChannelById)
		ChannelSubRoute.With(auth.JWT).Patch("/add/channelId={channelId}", controller.AddChannelMember)
		ChannelSubRoute.With(auth.JWT).Delete("/delete/channelId={channelId}", controller.DeleteChannelMember)
		ChannelSubRoute.Handle("/chat", http.HandlerFunc(controller.HandlerChannelWebSocket))
		ChannelSubRoute.With(auth.JWT).Patch("/addTaskColumn/channelId={channelId}", controller.AddTaskColumn)
		ChannelSubRoute.With(auth.JWT).Delete("/deleteTaskColumn/channelId={channelId}", controller.DeleteTaskColumn)
	})
}
