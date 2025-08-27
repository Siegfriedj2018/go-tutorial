package commands

import (
	"fmt"

	"go-tutorial/bootdotdev/pokedexcli/internal"
)

func CommandMapb(conf *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	
	
	response, err := internal.Connection(url)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for _, result := range response.Results {
		fmt.Println("\t", result.Name)
	}

	conf.Next = response.Previous
	conf.Previous = response.Previous
	return nil
}
