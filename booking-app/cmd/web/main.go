package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mkalashnikau/golang_app/booking-app/pkg/config"
	"github.com/mkalashnikau/golang_app/booking-app/pkg/handlers"
	"github.com/mkalashnikau/golang_app/booking-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// we write the template cache to the app variable
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
