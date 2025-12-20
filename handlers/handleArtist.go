package handlers

import (
	"net/http"
)

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func HandleArtist(w http.ResponseWriter, r any) {
	if ErrParse != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var arrArtist []artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := GetJson(url, &arrArtist)
	if err != nil {
		HandlerErr(w, "fetching Errore", http.StatusNotFound)
		return
	}
	Temp.ExecuteTemplate(w, "index.html", arrArtist)
}
