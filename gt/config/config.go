package config

import "gt/modals"

const (
	Port         = ":3000"
	APIartists   = "https://groupietrackers.herokuapp.com/api/artists"
	APIrelations = "https://groupietrackers.herokuapp.com/api/relation"
	APIlocation = "https://groupietrackers.herokuapp.com/api/locations"
	APIdates = "https://groupietrackers.herokuapp.com/api/dates"
)

var Fulldata []modals.FullArtistInfo
