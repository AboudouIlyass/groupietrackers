package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SolveFetchJSON() {
	FetchArtists, err := FetchArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(FetchArtists[0])
}

func FetchArtists(url string) ([]Artist, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists, err
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
}
