package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// main fetched function
func Fetch() ([]ArtistInfo, error) {
	fetchedArtists, errA := FetchArtists("https://groupietrackers.herokuapp.com/api/artists")
	if errA != nil {
		return []ArtistInfo{}, errA
	}

	for i, artist := range fetchedArtists {
		// here i have for example the data contain 8 locations from index 1 to 8
		fetchedLocations, errL := FetchLocations(artist.Locations)
		if errL != nil {
			return []ArtistInfo{}, errL
		}
		_, err := FetchDates(fetchedLocations[i].Dates)
		if err != nil {
			return nil, err
		}
	}

	return fetchedArtists, nil
}

// fetch artists
func FetchArtists(urlStr string) ([]ArtistInfo, error) {
	resp, errGet := http.Get(urlStr)
	if errGet != nil {
		return []ArtistInfo{}, errGet
	}
	defer resp.Body.Close()

	data, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return []ArtistInfo{}, errR
	}

	var Arists []ArtistInfo
	json.Unmarshal(data, &Arists)
	return Arists, nil
}

// fetch locations
func FetchLocations(url string) ([]LOCATIONS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []LOCATIONS{}, err
	}
	defer resp.Body.Close()

	data, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return []LOCATIONS{}, errR
	}
	var Locations []LOCATIONS
	json.Unmarshal(data, &Locations)

	return Locations, nil
}

// fetch dates
func FetchDates(url string) ([]DATES, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []DATES{}, err
	}
	defer resp.Body.Close()

	data, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return []DATES{}, errR
	}
	var dates []DATES
	json.Unmarshal(data, &dates)

	return dates, nil
}
