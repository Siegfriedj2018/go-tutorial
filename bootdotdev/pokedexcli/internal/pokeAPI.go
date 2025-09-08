package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func Connection(url string, cache *Cache) (*LocationArea, error) {
	body, ok := cache.Get(url)

	if !ok {
		log.Println("Location not in cache, connecting to Pokedex...")

		res, err := http.Get(url)
		if err != nil || res.StatusCode > 299 {
			return &LocationArea{}, fmt.Errorf("status code: %d\nerror: %w", res.StatusCode, err)
		}
		defer res.Body.Close()
		
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return &LocationArea{}, fmt.Errorf("error reading the body:: %w", err)
		}
		cache.Add(url, body)
	}

	var locations *LocationArea
	if err := json.Unmarshal(body, &locations); err != nil {
		return &LocationArea{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	
	return locations, nil
}

func RetrievePokemon() (*EncounterMethod, error) {
	
}