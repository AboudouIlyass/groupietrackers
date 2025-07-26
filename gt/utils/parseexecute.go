package utils

import (
	"html/template"
	"log"
	"net/http"
)

// parse and execute the templates
func ParseAndExecute(w http.ResponseWriter, filename string, artists interface{}) {
	tmp, err := template.ParseFiles(filename)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		log.Fatal(err)
	}
	err = tmp.Execute(w, artists)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		log.Fatal(err)
	}
}
