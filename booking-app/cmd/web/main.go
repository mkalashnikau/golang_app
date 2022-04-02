package main

import (
	"fmt"
	"net/http"

	"github.com/mkalashnikau/golang_app/booking-app/pkg/handlers"
	//"github.com/mkalashnikau/golang_udemy/building-web-app/4-reorganize-code/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
