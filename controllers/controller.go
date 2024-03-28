package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/teelowe/shall_we_sail/models"
	"github.com/teelowe/shall_we_sail/views"
)

func GetWeatherTidesHandler(w http.ResponseWriter, r *http.Request) {
	weatherResponseCh := make(chan models.Forecast)
	tidesResponseCh := make(chan models.Tides)

	go makeWeatherRequest(weatherResponseCh)
	go makeTidesRequest(tidesResponseCh)

	weatherResponse := <-weatherResponseCh
	tidesResponse := <-tidesResponseCh
	// log.Printf("tideResponse: %v", tidesResponse)

	unified := models.UnifiedWeatherAndTides{
		Forecasts: []models.ForecastOrTides{
			{Forecast: weatherResponse},
			{Tides: tidesResponse},
		},
	}
	// log.Printf("unified: %v", unified.Forecasts[1])
	views.RenderForecast(w, unified)
}

func makeWeatherRequest(ch chan<- models.Forecast) {
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
	ch <- forecasts
}

func makeTidesRequest(ch chan<- models.Tides) {
	now := time.Now().Format("20060102")
	oneWeek := time.Now().AddDate(0, 0, 6).Format("20060102")
	tidesUrl := fmt.Sprintf("https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?begin_date=%v&end_date=%v&station=9415339&product=predictions&datum=MLLW&time_zone=lst_ldt&interval=hilo&units=english&application=DataAPI_Sample&format=json", now, oneWeek)
	// log.Println("debug: " + tidesUrl)
	predictions, err := http.Get(tidesUrl)
	if err != nil {
		panic(err)
	}
	defer predictions.Body.Close()
	body, err := io.ReadAll(predictions.Body)
	if err != nil {
		panic(err)
	}

	var tides models.Tides
	if err = json.Unmarshal(body, &tides); err != nil {
		panic(err)
	}
	// log.Println(tides)
	ch <- tides
}
