package handlers

import (
	"net/http"

	"gt/utils"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		utils.ErrorPage(w, http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	utils.ParseAndExecute(w, "templates/html/artistsinfo.html")
}
