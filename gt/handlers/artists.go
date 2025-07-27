package handlers

import (
	"net/http"

	"gt/config"
	"gt/utils"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	utils.ParseAndExecute(w, "templates/html/artists.html", config.Fulldata)
}
