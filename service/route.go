package service

import (
	"web-service/pkg/router"
	"web-service/service/account"
	"web-service/service/channel"
	"web-service/service/index"
	"web-service/service/mail"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

	// Load account sub routes
	router.Router.Mount(router.RouterBasePath+"/account", account.AccountSubRoute)
	router.Router.Mount(router.RouterBasePath+"/mail", mail.MailSubRoute)
	router.Router.Mount(router.RouterBasePath+"/channel", channel.ChannelSubRoute)
}
