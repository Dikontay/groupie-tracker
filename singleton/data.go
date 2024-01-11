package singleton

import (
	"encoding/json"
	"fmt"
	"groupie-trakcer/models"

	"log"
	"net/http"
	"sync"
)

type data struct {
	AllArtists   []models.Artist
	AllDates     []models.Dates
	AllLocations []models.Locations
	AllRelations []models.Relations
}

type sinleInstance struct{
	Dates     models.Dates
	Locations models.Locations
	Relations models.Relations
}

var lock = &sync.Mutex{}

var (
	artistUrl = "https://groupietrackers.herokuapp.com/api/artists"
)
var singleData *data

func GetAllData() (*data, error) {
	if singleData == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleData == nil {
			var err error
			singleData, err = newData()
			if err != nil {
				return nil, err
			}
		}
	}
	return singleData, nil
}

func newData() (*data, error) {
	var artists []models.Artist

	err := getElement(artistUrl, &artists)
	if err != nil {
		log.Print("get element error")
		return nil, err
	}
	var wg sync.WaitGroup

	dataCh := make(chan *sinleInstance, len(artists))
	errCh := make(chan error, len(artists))

	for _, artist := range artists{
		wg.Add(1)
		go func (artist models.Artist){
			defer wg.Done()
			// single date, location, relation for the one artist
			var data sinleInstance
			var err error

			err = getElement(artist.ConcertDates, &data.Dates)
			if err != nil {
				errCh<-err
				return
			}
			
			err = getElement(artist.Locations, &data.Locations)
			if err != nil {
				errCh<-err
				return
			}
			err = getElement(artist.Relations, &data.Relations)
			if err != nil {
				errCh<-err
				return
			}

			dataCh<-&data


		}(artist)
	}
	 
	wg.Wait()
	close(dataCh)
	close(errCh)
	allData := &data{}

	allData.AllArtists=artists
	for d := range dataCh{
		allData.AllDates = append(allData.AllDates, d.Dates)
		allData.AllLocations=append(allData.AllLocations, d.Locations)
		allData.AllRelations=append(allData.AllRelations, d.Relations)
	}
	return allData, nil
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

