package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/teelowe/shall_we_sail/models"
	"github.com/teelowe/shall_we_sail/views"
)

func GetWeatherAndTidesHandler(w http.ResponseWriter, r *http.Request) {
	weather, err := http.Get("https://api.weather.gov/gridpoints/MTR/72,126/forecast")
	if err != nil {
		panic(err)
	}
	defer weather.Body.Close()
	body, err := io.ReadAll(weather.Body)
	if err != nil {
		panic(err)
	}

	var forecasts models.Forecast
	if err = json.Unmarshal(body, &forecasts); err != nil {
		panic(err)
	}
	views.RenderForecast(w, forecasts)
}
