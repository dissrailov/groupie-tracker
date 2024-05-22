package internal

import (
	"encoding/json"
	"errors"
	"groupie-tracker/structs"
	"io/ioutil"
	"net/http"
)

func GetLocation(url string) (structs.Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return structs.Location{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return structs.Location{}, err
	}
	// fmt.Println(body)
	var loca structs.Location
	if err := json.Unmarshal(body, &loca); err != nil {
		return structs.Location{}, err
	}
	return loca, nil
}

func GetData(url string) ([]structs.Artist, error) {
	res, err := http.Get(url)
	if err != nil {
		return []structs.Artist{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []structs.Artist{}, err
	}

	var artists []structs.Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return []structs.Artist{}, err
	}

	if len(artists) == 0 {
		return []structs.Artist{}, errors.New("no artists found")
	}

	return artists, nil
}

func GetRelation(url string) (structs.Relations, error) {
	res, err := http.Get(url)
	if err != nil {
		return structs.Relations{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return structs.Relations{}, err
	}
	// fmt.Println(body)
	var relation structs.Relations
	if err := json.Unmarshal(body, &relation); err != nil {
		return structs.Relations{}, err
	}
	return relation, nil
}
