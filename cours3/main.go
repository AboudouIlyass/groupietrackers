package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relations struct {
	Index []IndexItem `json:"index"`
}

type IndexItem struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func main() {
	// get the data
	respRel, errR := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if errR != nil {
		return
	}
	defer respRel.Body.Close()
	dataRel, _ := io.ReadAll(respRel.Body)

	var relations Relations
	if err := json.Unmarshal(dataRel, &relations); err != nil {
		return
	}

	respA, errA := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if errA != nil {
		return
	}
	defer respA.Body.Close()

	dataA, _ := io.ReadAll(respA.Body)

	var Artists Artist
	if err := json.Unmarshal(dataA, &Artists); err != nil {
		return
	}

	for i, v := range Artists.ID {
		if i == 1 {
			fmt.Println(v)
		}
	}
}
