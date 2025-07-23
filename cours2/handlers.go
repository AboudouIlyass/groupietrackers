package main

import (
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	if r.Method != http.MethodGet {
		return
	}
	fetcheddata, errF := Fetch("https://groupietrackers.herokuapp.com/api/artists")
	if errF != nil {
		return
	}

	tmp, errP := template.ParseFiles("html.html")
	if errP != nil {
		return
	}
	tmp.Execute(w, fetcheddata)
}
