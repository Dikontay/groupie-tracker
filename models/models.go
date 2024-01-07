package models

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbumDate string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"concertDates"`
	Relations      string   `json:"relations"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Locations struct {
	Index []struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
		DatesUrl string   `json:"dates"`
	} `json:"index"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
