package server

import (
	"fmt"
	"log"
	"net/http"

	"gt/handlers"
)

func Server() {
	// fitchedData, err := fetch.Fetch()
	// if err != nil {
	// 	log.Fatal("Error at fetching from api", err)
	// }
	fmt.Println("data is fitched ✅")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artists", handlers.Artists)
	mux.HandleFunc("/artistsinfo", handlers.ArtistsInfo)
	handlers.ServeAssetsAndPicsFiles(mux)

	s := &http.Server{
		Handler: mux,
		Addr:    ":3000",
	}
	log.Println("Server is running at http://localhost:3000")
	log.Fatal(s.ListenAndServe())
}
