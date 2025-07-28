package server

import (
	"log"
	"net/http"

	"gt/config"
	"gt/fetch"
	"gt/handlers"
)

func Server() {
	// fetch the data and store it in a globale variable
	fetch.SolveFetch()

	// handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artists", handlers.Artists)
	mux.HandleFunc("/artistsinfo/", handlers.ArtistsInfo)
	handlers.ServeAssetsAndPicsFiles(mux)

	s := &http.Server{
		Handler: mux,
		Addr:    config.Port,
	}
	log.Println("Server is running at http://localhost:3000")

	log.Fatal(s.ListenAndServe())
}

/*

<body>
    {{range .}}
        <img src="{{.Image}}" />
        <p>Name: {{.Name}}</p>

        {{range .Relations.Index}}
            {{range $loc, $dates := .DatesLocations}}
                    {{$loc}} :
                    {{range $dates}}
                       <p> {{.}}</p>
                    {{end}}
            {{end}}
        {{end}}
    {{end}}
</body>



*/
