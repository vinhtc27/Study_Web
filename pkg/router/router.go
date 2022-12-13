package router

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	// "web-service/pkg/cache"

	"web-service/pkg/db"
	"web-service/pkg/log"
	"web-service/pkg/server"
)

// RouterBasePath Variable
var RouterBasePath string

// Router Variable
var Router *chi.Mux

// Initialize Function in Router
func init() {
	// Initialize Router
	Router = chi.NewRouter()
	Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Set Router Handler
	Router.NotFound(handlerNotFound)
	Router.MethodNotAllowed(handlerMethodNotAllowed)
	Router.Get("/favicon.ico", handlerFavIcon)
}

// HealthCheck Function
func HealthCheck(w http.ResponseWriter) {
	// Check Database Connections
	if len(server.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
		case "postgres":
			err := db.PSQL.Ping()
			if err != nil {
				log.Println(log.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		}
	}

	// Return Success response
	ResponseSuccess(w, "200", "")
}
