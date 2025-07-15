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

	tmp, er := template.ParseFiles("static/pages/home.html")
	if er != nil {
		log.Println("Error parsing home.html file", er)
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	FetchedData, errr := fetchdata.Fetch("https://fakestoreapi.com/products/")
	if errr != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	err := tmp.Execute(w, FetchedData)
	if err != nil {
		log.Println("error during execution")
		utils.RenderError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pages/css/home.css")
}
