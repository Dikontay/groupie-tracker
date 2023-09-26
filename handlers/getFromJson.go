package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

func getElement(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return err
	}
	return nil
}


