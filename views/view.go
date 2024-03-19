package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/teelowe/shall_we_sail/models"
)

func RenderForecast(w http.ResponseWriter, data models.Forecast) {
	tmpl, err := template.ParseFiles("templates/template.gohtml")
	if err != nil {
		http.Error(w, "Parse error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
