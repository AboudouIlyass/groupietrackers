package main

type ArtistInfo struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
}

type LOCATIONS struct {
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DATES struct {
	Dates []string `json:"dates"`
}
