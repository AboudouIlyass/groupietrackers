package server

import (
	"log"
	"net/http"

	"groupietrackers/handlers"
)

type Router struct{}

const (
	Port = ":8080"
)

func Server() {
	var router Router
	log.Println("Server is running at http://localhost" + Port)
	er := http.ListenAndServe(Port, router)
	if er != nil {
		log.Fatal(er)
	}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.HomeHandler(w, r)
	case "/static/pages/css/home.css":
		handlers.StaticHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

