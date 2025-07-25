package utils

import (
	"html/template"
	"net/http"

	"gt/modals"
)

// error page rendred
func ErrorPage(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("templates/html/error.html")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	switch status {
	case 400:
		w.WriteHeader(status)
		err = tmp.Execute(w, modals.Errpage{Error: status, Msg: "Bad Request"})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	case 404:
		w.WriteHeader(status)
		err = tmp.Execute(w, modals.Errpage{Error: status, Msg: "Page Not Found"})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	case 405:
		w.WriteHeader(status)
		err = tmp.Execute(w, modals.Errpage{Error: status, Msg: "Method Not Allowed"})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	case 500:
		w.WriteHeader(status)
		err = tmp.Execute(w, modals.Errpage{Error: status, Msg: "Internal Server Error"})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
}
