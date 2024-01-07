package singleton

import (
	"encoding/json"
	"fmt"
	"groupie-trakcer/models"

	"log"
	"net/http"
	"sync"
)

type Data struct {
	AllArtists   []models.Artist
	AllDates     []models.Dates
	AllLocations []models.Locations
	AllRelations []models.Relations
}

var lock = &sync.Mutex{}

var (
	artistUrl = "https://groupietrackers.herokuapp.com/api/artists"
)
var singleData *Data

func GetAllData() *Data {
	if singleData == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleData == nil {
			fmt.Println("Creating single instance now.")
			singleData = newData()
		}
	}
	return singleData
}

func newData() *Data {
	var artists []models.Artist
	var dates []models.Dates
	var locations []models.Locations
	var relations []models.Relations
	err := getElement(artistUrl, &artists)
	if err != nil {
		log.Print("get element error")
		return nil
	}
	for _, artist := range artists {
		date := models.Dates{}
		err = getElement(artist.ConcertDates, &date)
		if err != nil {
			log.Printf("Error getting date for artist : %T", artist)
			return nil
		}
		dates = append(dates, date)

		location := models.Locations{}
		err = getElement(artist.Locations, &location)
		if err != nil {
			log.Printf("Error getting location for artist : %T", artist)
			return nil
		}
		locations = append(locations, location)

		relation := models.Relations{}
		err = getElement(artist.Relations, &relation)
		if err != nil {
			log.Printf("Error getting relation for artist : %T", artist)
			return nil
		}

		relations = append(relations, relation)
	}

	singleData = &Data{AllArtists: artists, AllDates: dates, AllLocations: locations, AllRelations: relations}

	return singleData
}

func getElement(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}
