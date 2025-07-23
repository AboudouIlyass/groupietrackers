package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
