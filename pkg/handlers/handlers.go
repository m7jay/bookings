package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/m7jay/bookings/pkg/config"
	"github.com/m7jay/bookings/pkg/models"
	"github.com/m7jay/bookings/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("request to / with request header %v", middleware.GetReqID(r.Context()))
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Printf("request to /about with requset header %v", middleware.GetReqID(r.Context()))
	stringMap := make(map[string]string)
	stringMap["test"] = "hello!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.Template(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
