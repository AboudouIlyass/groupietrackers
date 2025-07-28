package fetch

import (
	"log"

	"gt/config"
	"gt/modals"
)

func SolveFetch() {
	var err error
	// fetch artists
	var artists []modals.Artist
	err = Fetch(config.APIartists, &artists)
	if err != nil {
		log.Println(err)
		return
	}
	// fetch relations
	var relations modals.Relations
	err = Fetch(config.APIrelations, &relations)
	if err != nil {
		log.Println(err)
		return
	}

	// fetch locations
	var locLocations modals.LOCATIONS
	err = Fetch(config.APIlocation, &locLocations)
	if err != nil {
		log.Println(err)
		return
	}

	// combine them into 'config.Fulldata'
	for _, art := range artists {
		var rel modals.Relations
		for _, r := range relations.Index {
			if art.ID == r.ID {
				rel.Index = append(rel.Index, r)
				break
			}
		var loc modals.LOCATIONS
		for _, l := range locLocations.Index{
			if l.ID == art.ID{
				loc.Index = append(loc.Index, l)
				break
			}
	
		}
		}
		config.Fulldata = append(config.Fulldata, modals.FullArtistInfo{
			Artist:    art,
			Relations: rel,
			LOCATIONS: locLocations,
		})
	}
}
