package modals

// the data contain all the informations
type ArtistFullInfo struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Artist struct {
	ID           int         `json:"id"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate int         `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
	LocationsURL string `json:"locations"`
	DatesURL     string      `json:"concertDates"`
	RelationsURL string      `json:"relations"`
}

type Locations struct {
	Loc string `json:"locations"`
}

type Date struct {
	Dates string `json:"dates"`
}
