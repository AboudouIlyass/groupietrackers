package handlers

import "net/http"

func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	utils.ParseAndExecute(w, "templates/html/artistsinfo.html", config.Fulldata)
}
