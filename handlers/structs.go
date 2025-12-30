package handlers

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
	ID        int      `json:"id"`
	Dateslist []string `json:"dates"`
}
type Relations struct {
	ID       int                 `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

type InfosPage struct {
	ArtitsInfo    Artist
	LocationsInfo Locations
	DatesInfo     Dates
	Relationsinfo Relations
}
