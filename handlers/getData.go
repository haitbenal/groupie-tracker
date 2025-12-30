package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func FetchData() []Artist {
	Urlartist := "https://groupietrackers.herokuapp.com/api/artists"
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

func FetchRelations(url string) *Relations {
	var relation Relations
		res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&relation)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &relation
}