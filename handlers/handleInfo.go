package handlers

import (
	"net/http"
)

type LocationStruct struct {
	ID            int      `json:"id"`
	LocationArr   []string `json:"locations"`
	LocationDates string   `json:"dates"`
}
type ConcertStruct struct {
	ID         int      `json:"id"`
	ConcertArr []string `json:"dates"`
}
type RelationsStruct struct {
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	HandleLocations(w, r)

}
