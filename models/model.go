package models

import (
	"log"
	"time"
)

type Forecast struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Periods []Period `json:"periods"`
}

type Period struct {
	Number           int64     `json:"number"`
	Name             string    `json:"name"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	IsDaytime        bool      `json:"isDaytime"`
	Temperature      int64     `json:"temperature"`
	TemperatureUnit  string    `json:"temperatureUnit"`
	TemperatureTrend string    `json:"temperatureTrend"`
	WindSpeed        string    `json:"windSpeed"`
	WindDirection    string    `json:"windDirection"`
	ShortForecast    string    `json:"shortForecast"`
	DetailedForecast string    `json:"detailedForecast"`
}

func FormatDate(t time.Time) string {
	return t.Format("Jan 2")
}

func FormatTidesDate(s string) string {
	t, err := time.Parse("2006-01-02 15:04", s)
	if err != nil {
		log.Fatal("Date format error")
	}
	return t.Format("Jan 2, 15:04")
}

type Tides struct {
	Predictions []Datum `json:"predictions"`
}

type Datum struct {
	Time  string `json:"t"`
	Value string `json:"v"`
	HiLo  string `json:"type"`
}

type UnifiedWeatherAndTides struct {
	Forecasts []ForecastOrTides
}

type ForecastOrTides struct {
	Forecast Forecast
	Tides    Tides
}
