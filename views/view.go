package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/teelowe/shall_we_sail/models"
)

func RenderForecast(w http.ResponseWriter, forecast models.UnifiedWeatherAndTides) {
	tmpl, err := template.New("template.gohtml").Funcs(template.FuncMap{
		"FormatForecastDate": models.FormatForecastDate,
		"FormatTidesDate":    models.FormatTidesDate}).ParseFiles("templates/template.gohtml")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Parse error: %v", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, forecast)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Execute error: %v", http.StatusInternalServerError)
		return
	}
}
