package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Members []string `json:"members"`
	Data    int      `json:"creationDate"`
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