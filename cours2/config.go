package main

type ArtistInfo struct {
	ID           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      []string `json:"members"`
	CreationDate int    `json:"creationDate"`
}
