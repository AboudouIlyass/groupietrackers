package handlers

import "net/http"

func ServeAssetsAndPicsFiles(mux *http.ServeMux) {
	mux.HandleFunc("/templates/assets/error.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/assets/error.css")
	})
	mux.HandleFunc("/templates/assets/home.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/assets/home.css")
	})
	
	mux.HandleFunc("/templates/pics/logo.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/pics/logo.png")
	})
	mux.HandleFunc("/templates/pics/b.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/pics/b.jpg")
	})
}
