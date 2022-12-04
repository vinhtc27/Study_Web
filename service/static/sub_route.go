package static

import (
	"net/http"
	"path"

	"github.com/go-chi/chi"
)

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
	http.ServeFile(w, r, "./build/index.html")
}

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	cssFileName := path.Base(r.URL.Path)
	http.ServeFile(w, r, "./build/static/css/"+cssFileName)
}

func ServeJS(w http.ResponseWriter, r *http.Request) {
	jsFileName := path.Base(r.URL.Path)
	http.ServeFile(w, r, "./build/static/js/"+jsFileName)
}
