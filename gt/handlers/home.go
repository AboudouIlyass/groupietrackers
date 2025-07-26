package handlers

import (
	"net/http"

	"gt/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorPage(w, http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorPage(w, http.StatusBadRequest)
		return
	}
	utils.ParseAndExecute(w, "templates/html/home.html", nil)
}
