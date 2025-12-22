package handlers

import (
	"fmt"
	"net/http"
)

func HandleLocations(w http.ResponseWriter, r *http.Request) {
	var locatonsstruct LocationStruct
	fmt.Println(r.URL.Path)
	id := r.URL.Path[len("/artist/"):]
	locations := "https://groupietrackers.herokuapp.com/api/locations/" + id
	err := GetJson(locations, &locatonsstruct)
	if err != nil {
		HandlerErr(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Println(locatonsstruct.ID)
	fmt.Println(locatonsstruct.LocationArr)
	fmt.Println(locatonsstruct.LocationDates)

	// concertDates := "https://groupietrackers.herokuapp.com/api/dates/" + id
	// relations := "https://groupietrackers.herokuapp.com/api/relation/" + id
	// fmt.Println(locations, concertDates, relations)
}
