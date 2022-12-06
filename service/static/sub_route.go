package static

import (
	"net/http"
	"path"
	"web-service/pkg/server"

	"github.com/go-chi/chi"
)

var fePath = server.Config.GetString("FE_PATH")

var StaticSubRoute = chi.NewRouter()

// Init package with sub-router for mails service
func init() {
	StaticSubRoute.Group(func(_ chi.Router) {
		StaticSubRoute.Get("/html", ServeHTML)
		StaticSubRoute.Get("/css/*", ServeCSS)
		StaticSubRoute.Get("/js/*", ServeJS)
	})
}

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fePath+"/index.html")
}

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	cssFileName := path.Base(r.URL.Path)
	http.ServeFile(w, r, fePath+"/static/css/"+cssFileName)
}

func ServeJS(w http.ResponseWriter, r *http.Request) {
	jsFileName := path.Base(r.URL.Path)
	http.ServeFile(w, r, fePath+"/static/js/"+jsFileName)
}
