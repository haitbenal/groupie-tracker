package handlers

import (
	"net/http"
	"strconv"
)

func GetArtist(artists []Artist, id int) *Artist {
	for i := range artists {
		if artists[i].ID == id {
			return &artists[i]
		}
	}
	return nil
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	allartist := FetchData()
	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		HandelError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	data := GetArtist(allartist, userID)
	if data == nil {
		HandelError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	data_location := FetchLocations(data.Urllocations)
	if data_location == nil {
		HandelError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	data_dates := FetchDates(data.Urlconcertdates)
	if data_dates == nil {
		HandelError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	final := InfosPage{ArtitsInfo: *data, LocationsInfo: *data_location, DatesInfo: *data_dates}

	errpage := Temp.ExecuteTemplate(w, "artistInfo.html", final)
	if errpage != nil {
		HandelError(w, "hhh Server Error", http.StatusInternalServerError)
		return
	}
}
