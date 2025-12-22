package handlers

import (
	"fmt"
	"net/http"
)

type LocationStruct struct {
	Id       int      `json:"id"`
	Locatins []string `json:"locations"`
}
type DatesStruct struct {
	Id       int      `json:"id"`
	Locatins []string `json:"dates"`
}

type RelationStruct struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func handleLocations(w http.ResponseWriter, r *http.Request) {
	var locatonsstruct LocationStruct

	id := r.URL.Path[len("/artist/"):]
	locations := "https://groupietrackers.herokuapp.com/api/locations/" + id
	err := GetJson(locations, &locatonsstruct)
	if err != nil {
		HandlerErr(w, "Not Found", http.StatusNotFound)
		return
	}
	Temp.ExecuteTemplate(w, "artistInfo.html", locatonsstruct)

}
func handleDates(w http.ResponseWriter, r *http.Request) {
	var datesstruct DatesStruct
	id := r.URL.Path[len("/artist/"):]
	locations := "https://groupietrackers.herokuapp.com/api/dates/" + id
	err := GetJson(locations, &datesstruct)
	if err != nil {
		fmt.Println(err)
		HandlerErr(w, "Not Found", http.StatusNotFound)
		return
	}
	Temp.ExecuteTemplate(w, "artistInfo.html", datesstruct)

}

func handleRelation(w http.ResponseWriter, r *http.Request) {
	var relationstruct RelationStruct
	id := r.URL.Path[len("/artist/"):]
	locations := "https://groupietrackers.herokuapp.com/api/relation/" + id
	err := GetJson(locations, &relationstruct)
	if err != nil {
		fmt.Println(err)
		HandlerErr(w, "Not Found", http.StatusNotFound)
		return
	}
	Temp.ExecuteTemplate(w, "artistInfo.html", relationstruct)

}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	handleLocations(w, r)
	handleDates(w, r)
	handleRelation(w, r)

}
