package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func Fetch(urlStr string) ([]ArtistInfo, error) {
	// get the data
	resp, errGet := http.Get(urlStr)
	if errGet != nil {
		return []ArtistInfo{}, errGet
	}
	defer resp.Body.Close()

	// read the data
	data, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return []ArtistInfo{}, errR
	}

	//
	var Arists []ArtistInfo
	json.Unmarshal(data, &Arists)
	return Arists, nil
}
