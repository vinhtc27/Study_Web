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
	})
}
