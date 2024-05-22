package structs

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type AllArtist struct {
	Artist    Artist
	Relations Relations
	Location  Location
}
