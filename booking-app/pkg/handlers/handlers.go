package handlers

import (
	"net/http"

	"github.com/mkalashnikau/golang_app/booking-app/pkg/render"
	//"github.com/mkalashnikau/golang_udemy/building-web-app/4-reorganize-code/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
