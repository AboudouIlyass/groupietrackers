package modals

type FullArtistInfo struct {
	Artist
	Relations
	LOCATIONS
}

// artists
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// relations
type Relations struct {
	Index Index
}

type Index []struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}


// locations 
type LOCATIONS struct{
	Index Ind `json:"index"`
}

type Ind []struct	{
	ID int	`json:"id"`
	Locations []string `json:"locations"`
}