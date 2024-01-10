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

func GetAllData() (*Data, error) {
	if singleData == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleData == nil {
			fmt.Println("Creating single instance now.")
			var err error
			singleData, err = newData()
			if err != nil {
				return nil, err
			}
		}
	}
	return singleData, nil
}

func newData() (*Data, error) {
	var artists []models.Artist
	var dates []models.Dates
	var locations []models.Locations
	var relations []models.Relations
	

	var err, lastErr error

	err = getElement(artistUrl, &artists)
	if err != nil {
		log.Print("get element error")
		return nil, err
	}

	for _, artist := range artists {
		date := models.Dates{}
		err = getElement(artist.ConcertDates, &date)
		if err != nil {
			lastErr = err
			log.Printf("Error getting date for artist : %T", artist)
		}
		dates = append(dates, date)

		location := models.Locations{}
		err = getElement(artist.Locations, &location)
		if err != nil {
			log.Printf("Error getting location for artist : %T", artist)
			lastErr= err
			
		}
		locations = append(locations, location)

		relation := models.Relations{}
		err = getElement(artist.Relations, &relation)
		if err != nil {
			log.Printf("Error getting relation for artist : %T", artist)
			lastErr = err
		}

		relations = append(relations, relation)
	}
	if lastErr != nil {
		return nil, lastErr
	}

	 

	return &Data{AllArtists: artists, AllDates: dates, AllLocations: locations, AllRelations: relations}, nil
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

