package service

import (
	"web-service/pkg/router"
	"web-service/service/index"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

}
