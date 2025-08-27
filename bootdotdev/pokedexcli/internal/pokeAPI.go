package internal

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


func Connection(url string) (*LocationArea, error) {
	fmt.Println("Connecting to Pokedex...")
	res, err := http.Get(url)
	if err != nil || res.StatusCode > 299 {
		return &LocationArea{}, fmt.Errorf("status code: %d\nerror: %w", res.StatusCode, err)
	}
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &LocationArea{}, fmt.Errorf("error reading the body:: %w", err)
	}

	var locations *LocationArea
	if err := json.Unmarshal(body, &locations); err != nil {
		return &LocationArea{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	
	return locations, nil
}
