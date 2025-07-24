package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
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

// the full struct info
type ArtistsFullInfo struct {
	Artists       Artist
	LocationDates []LocationDates
}

type LocationDates struct {
	Location string
	Dates    []string
}

func main() {
	// fetch relations
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

	// fetch artists
	respA, errA := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if errA != nil {
		return
	}
	defer respA.Body.Close()

	dataA, _ := io.ReadAll(respA.Body)

	var Artists []Artist
	if err := json.Unmarshal(dataA, &Artists); err != nil {
		return
	}

	// build a slice combined the artists & relations
	

	/*
		 Build map from artist ID to IndexItem (relations)
			    relMap := make(map[int]IndexItem)
			    for _, rel := range relations.Index {
			        relMap[rel.ID] = rel
			    }

			 // Build combined slice
			    var fullInfos []ArtistsFullInfo
			    for _, artist := range artists {
			        var locationDates []LocationDates
			        if rel, ok := relMap[artist.ID]; ok {
			            for loc, dates := range rel.DatesLocations {
			                locationDates = append(locationDates, LocationDates{
			                    Location: loc,
			                    Dates:    dates,
			                })
			            }
			        }

			        fullInfos = append(fullInfos, ArtistsFullInfo{
			            Artists:       artist,
			            LocationDates: locationDates,
			        })
			    }

	*/

	// parse function
	parsehtmlfile := func(path string) (*template.Template, error) {
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return nil, err
		}
		return tmpl, nil
	}

	// hanlde the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := parsehtmlfile("h.html")
		if err != nil {
			http.Error(w, "Internal Server Error: failed to load template", http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case http.MethodGet:
			if err := tmp.Execute(w, Artists); err != nil {
				http.Error(w, "Internal Server Error: failed to render template", http.StatusInternalServerError)
				return
			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	})

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
