package internal

import (
	"bootdevproject/pokedexcli/commands"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func Connection() *commands.Config {
	log.Println("Connecting to Pokedex...")
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/?offset=0&limit=20")
	if err != nil || res.StatusCode > 299 {
		log.Fatal("status code: ", res.StatusCode, "\nerror: ", err)
	}
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var conf *commands.Config
	if err := json.Unmarshal(body, &conf); err != nil {
		log.Fatalf("Error unmarshalling json: %v", err)
	}
	
	return conf
}