package main

import (
	"log"
	"net/http"

	"github.com/teelowe/shall_we_sail/controllers"
)

func main() {
	http.HandleFunc("/shall_we_sail", controllers.GetWeatherTidesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
