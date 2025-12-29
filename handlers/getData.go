package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Image           string   `json:"image"`
	Members         []string `json:"members"`
	Data            int      `json:"creationDate"`
	Album           string   `json:"firstAlbum"`
	Urllocations    string   `json:"locations"`
	Urlconcertdates string   `json:"concertDates"`
	Urlrelations    string   `json:"relations"`
}
type Locations struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
}

type Dates struct {
	ID       int      `json:"id"`
	Dateslist []string `json:"dates"`
}

type InfosPage struct{
	ArtitsInfo Artist
	LocationsInfo Locations
	DatesInfo Dates
}

var Urlartist = "https://groupietrackers.herokuapp.com/api/artists"


func FetchData() []Artist {
	var artists []Artist
	res, err := http.Get(Urlartist)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return artists
}


func FetchLocations(url string) *Locations {
	var location Locations
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &location
}
func FetchDates(url string) *Dates {
	var dates Dates
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&dates)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &dates
}