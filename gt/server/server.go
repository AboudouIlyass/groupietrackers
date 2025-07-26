package server

import (
	"log"
	"net/http"

	"gt/config"
	"gt/handlers"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artists", handlers.Artists)
	mux.HandleFunc("/artistsinfo", handlers.ArtistsInfo)
	handlers.ServeAssetsAndPicsFiles(mux)

	s := &http.Server{
		Handler: mux,
		Addr:    config.Port,
	}
	log.Println("Server is running at http://localhost:3000")

	log.Fatal(s.ListenAndServe())
}
