package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"gt/config"
	"gt/utils"
)

func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

	StrId := strings.TrimPrefix(r.URL.Path, "/artistsinfo/")
	id, err := strconv.Atoi(StrId)
	if err != nil {
		utils.ErrorPage(w, http.StatusBadRequest)
		return
	}
	// check the id 
	for _, artist := range config.Fulldata {
		if artist.ID == id {
			utils.ParseAndExecute(w, "templates/html/artistsinfo.html", artist)
			return
		}
	}

	utils.ErrorPage(w, http.StatusNotFound)
}
