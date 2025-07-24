package main

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	if r.Method != http.MethodGet {
		return
	}
	fetcheddata, err := Fetch()
	if err != nil {
		return
	}

	tmp, errP := template.ParseFiles("html.html")
	if errP != nil {
		return
	}

	tmp.Execute(w, fetcheddata)
}
