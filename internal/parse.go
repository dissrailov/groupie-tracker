package internal

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structs"
	"net/http"
	"strconv"
)

type Artist struct {
	Relations string `json:"relations"`
}

const (
	ArtistUrl = "https://groupietrackers.herokuapp.com/api/artists"
	Locations = "https://groupietrackers.herokuapp.com/api/locations/"
	Dates     = "https://groupietrackers.herokuapp.com/api/dates/"
	Relation  = "https://groupietrackers.herokuapp.com/api/relation/"
)

func ParseApi(idString string) (structs.AllArtist, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return structs.AllArtist{}, err
	}
	datarelation, err := GetRelation(Relation + idString)
	if err != nil {
		return structs.AllArtist{}, err
	}
	data, err := GetData(ArtistUrl)
	if err != nil {
		return structs.AllArtist{}, err
	}
	datalocation, err := GetLocation(Locations + idString)
	if err != nil {
		return structs.AllArtist{}, err
	}
	AllArtist := structs.AllArtist{
		Artist:    data[id-1],
		Relations: datarelation,
		Location:  datalocation,
	}
	return AllArtist, nil
}

func IdLimit() int {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return 0
	}

	groupIDs := make(map[int]bool)

	for _, artist := range artists {

		id := extractIDFromURL(artist.Relations)
		groupIDs[id] = true
	}

	numGroupIDs := len(groupIDs)
	return numGroupIDs
}

func extractIDFromURL(url string) int {
	var id int
	fmt.Sscanf(url, "https://groupietrackers.herokuapp.com/api/relation/%d", &id)
	return id
}
