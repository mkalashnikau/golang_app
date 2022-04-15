package handlers

import (
	"net/http"

	"github.com/mkalashnikau/golang_app/booking-app/pkg/config"
	"github.com/mkalashnikau/golang_app/booking-app/pkg/render"
)

// Repo is the repository used by the handlers (actually this is 'app' variable)
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

// NewHandlers sets the repository for the hadlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
// Thanks to (m *Repository) the handlers have access to the Repository variable. Also it is needed to specify Repo in the call in main.go: http.HandleFunc("/", handlers.Repo.Home)
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
