package handlers

import (
	"net/http"

	"github.com/toastsandwich/bookings/pkg/config"
	"github.com/toastsandwich/bookings/pkg/models"
	"github.com/toastsandwich/bookings/pkg/render"
)

// Repository struct and variable
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (rr *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rr.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.gohtml", &models.TemplateData{})
}

func (rr *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"
	remoteIp := rr.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (rr *Repository) WhoAmI(w http.ResponseWriter, r *http.Request) {
	remoteIp := rr.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "whoami.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
