package handlers

import (
	"log"
	"net/http"

	"gt/config"
	"gt/fetch"
	"gt/modals"
	"gt/utils"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method != http.MethodGet {
		utils.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	// fetch and render the data
	var artists []modals.Artist
	err = fetch.FetchArtists(config.APIartists, &artists)
	if err != nil {
		log.Println(err)
		return
	}
	var fullartistinfo []modals.ArtistFullInfo

	for _, artist := range artists {

		// fetch locations
		var loc modals.Locations
		err = fetch.FetchArtists(artist.LocationsURL, &loc)
		if err != nil {
			log.Println(err)
			continue
		}

		// fetch dates
		var d modals.Date
		err = fetch.FetchArtists(artist.LocationsURL, &d)
		if err != nil {
			log.Println(err)
			continue
		}

		// 

	}

	utils.ParseAndExecute(w, "templates/html/artists.html", artists)
}
