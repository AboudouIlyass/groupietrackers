package fetchdata

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"groupietrackers/config"
)

func Fetch(dataLink string) ([]config.Data, error) {
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
	var products []config.Data
	json.Unmarshal([]byte(body), &products)
	return products, nil
}
