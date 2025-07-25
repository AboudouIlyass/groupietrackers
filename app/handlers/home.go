package handlers

import (
	"log"
	"net/http"
	"text/template"

	"groupietrackers/fetchdata"
	"groupietrackers/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.RenderError(w, http.StatusNotFound, "page not found")
		return
	}
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	tmp, er := template.ParseFiles("static/html/home.html")
	if er != nil {
		log.Println("Error parsing home.html file", er)
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	fetcheddata, err := fetchdata.Fetch()
	if err != nil{
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	// execute
	err = tmp.Execute(w, fetcheddata)
	if err != nil {
		log.Println("error during execution")
		utils.RenderError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/css/home.css")
}

func PicturesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pics/gt.png")
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/js/popup.js")
}
