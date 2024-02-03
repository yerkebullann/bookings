package handlers

import (
	"MyFirstApp/pkg/config"
	"MyFirstApp/pkg/models"
	"MyFirstApp/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRep(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandles(r *Repository) {
	Repo = r
}
func (rs *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rs.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RrTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (rs *Repository) About(wrtAbout http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test": "Hello it is from Map",
	}
	remoteIP := rs.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	render.RrTemplate(wrtAbout, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
