package fetchdata

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"groupietrackers/config"
)
// fetch the artists json file
func FetchArt(dataLink string) ([]config.ArtistInfo, error) {
	response, err := http.Get(dataLink)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer response.Body.Close()

	body, er := io.ReadAll(response.Body)
	if er != nil {
		log.Fatal(er)
		return nil, er
	}
	var products []config.ArtistInfo
	json.Unmarshal([]byte(body), &products)
	return products, nil
}

// fetch the relations json file
func FetchRel(dataLink string) (config.Relations, error) {
	response, err := http.Get(dataLink)
	if err != nil {
		log.Fatal(err)
		return config.Relations{}, err
	}
	defer response.Body.Close()

	body, er := io.ReadAll(response.Body)
	if er != nil {
		log.Fatal(er)
		return config.Relations{}, er
	}
	var relations config.Relations
	json.Unmarshal([]byte(body), &relations)
	return relations, nil
}
