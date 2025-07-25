package modals

// the data contain all the informations
type ArtistFullInfo struct {
	ID            int
	Image         string
	Name          string
	Members       []string
	CreationDate  int
	FirstAlbum    string
	Locations     string
	ConcertDates  string
	Relations     string

	//LocationDates []LocationDates
}
