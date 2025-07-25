package config

// for artists
type ArtistInfo struct {
	ID           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      []string `json:"members"`
	CreationDate int    `json:"creationDate"`
}

// for relations
type Relations struct {
	Index []IndexItem `json:"index"`
}

type IndexItem struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// combine artists-relatoins
type ArtistFullInfo struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	LocationDates []LocationDates
}

type LocationDates struct {
	Location string
	Dates    []string
}