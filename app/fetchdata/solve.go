package fetchdata

import "groupietrackers/config"

func Fetch() ([]config.ArtistFullInfo,error) {
	// fetch artists
	FetchedArtists, err := FetchArt("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil,err
	}
	// fetch relations
	Fetchedrelations, err := FetchRel("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil,err
	}
	// merge artists and relations
	relationMap := map[int]map[string][]string{}
	for _, v := range Fetchedrelations.Index {
		relationMap[v.ID] = v.DatesLocations
	}

	var fullData []config.ArtistFullInfo
	for _, v := range FetchedArtists {
		var LocD []config.LocationDates
		if DL, exist := relationMap[v.ID]; exist {
			for l, dates := range DL {
				LocD = append(LocD, config.LocationDates{Location: l, Dates: dates})
			}
		}
		fullData = append(fullData, config.ArtistFullInfo{ID: v.ID, Image: v.Image, Name: v.Name, Members: v.Members, CreationDate: v.CreationDate, LocationDates: LocD})
	}
	return fullData, nil
}
