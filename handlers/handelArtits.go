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
	errpage := Temp.ExecuteTemplate(w, "artistInfo.html", data)
	if errpage != nil {
		HandelError(w, "hhh Server Error", http.StatusInternalServerError)
		return
	}
}
