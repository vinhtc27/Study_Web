package service

import (
	"web-service/pkg/router"
	"web-service/service/account"
	"web-service/service/channel"
	"web-service/service/mail"
	"web-service/service/static"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Load account sub routes
	router.Router.Mount(router.RouterBasePath+"/", static.StaticSubRoute)
	router.Router.Mount(router.RouterBasePath+"/account", account.AccountSubRoute)
	router.Router.Mount(router.RouterBasePath+"/mail", mail.MailSubRoute)
	router.Router.Mount(router.RouterBasePath+"/channel", channel.ChannelSubRoute)
}
