// package main

// import (
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", Home)

// 	log.Println("Server is running at http://localhost:3000")
// 	http.ListenAndServe(":3000", nil)
// }
// var pageTemplate = template.Must(template.ParseFiles("html.html"))

package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
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

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

type LocationDates struct {
	Location string
	Dates    []string
}

type ArtistFullInfo struct {
	Artist        Artist
	LocationDates []LocationDates
}

var pageTemplate = template.Must(template.New("page").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Groupie Trackers</title>
  <style>
    body { font-family: sans-serif; background: #f0f0f0; }
    .container { display: flex; flex-wrap: wrap; justify-content: center; }
    .card { background: white; margin: 10px; padding: 15px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1); width: 300px; }
    .card img { width: 100%; border-radius: 4px; }
    ul { list-style: none; padding: 0; margin: 5px 0; }
  </style>
</head>
<body>
  <h1 style="text-align:center;">Groupie Trackers</h1>
  <div class="container">
    {{range .}}
    <div class="card">
      <img src="{{.Artist.Image}}" alt="{{.Artist.Name}}">
      <h2>{{.Artist.Name}}</h2>
      <p><strong>Started:</strong> {{.Artist.CreationDate}}</p>
      <p><strong>First Album:</strong> {{.Artist.FirstAlbum}}</p>
      <h3>Members</h3>
      <ul>{{range .Artist.Members}}<li>{{.}}</li>{{end}}</ul>
      <h3>Concerts</h3>
      <ul>
        {{range .LocationDates}}
        <li><strong>{{.Location}}</strong>
          <ul>{{range .Dates}}<li>{{.}}</li>{{end}}</ul>
        </li>
        {{end}}
      </ul>
    </div>
    {{end}}
  </div>
</body>
</html>`))

func fetchArtists(url string) ([]Artist, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var artists []Artist
	if err := json.Unmarshal(data, &artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func fetchRelations(url string) ([]Relation, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var idx RelationIndex
	if err := json.Unmarshal(data, &idx); err != nil {
		return nil, err
	}
	return idx.Index, nil
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Error loading artists", 500)
		return
	}
	relations, err := fetchRelations("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		http.Error(w, "Error loading relations", 500)
		return
	}

	relMap := make(map[int]Relation, len(relations))
	for _, rel := range relations {
		relMap[rel.ID] = rel
	}

	var infos []ArtistFullInfo
	for _, art := range artists {
		locs := relMap[art.ID].DatesLocations
		var ld []LocationDates
		for loc, dates := range locs {
			ld = append(ld, LocationDates{loc, dates})
		}
		infos = append(infos, ArtistFullInfo{Artist: art, LocationDates: ld})
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pageTemplate.Execute(w, infos); err != nil {
		log.Printf("Template error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handlerHome)
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
