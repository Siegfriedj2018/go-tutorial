package commands

import (
	"fmt"

	"go-tutorial/bootdotdev/pokedexcli/internal"
)

func CommandMap(conf *Config, cache *internal.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if conf.Next != nil {
		url = *conf.Next
	} else {
		fmt.Println("You have reached the end of the map locations, starting over...")
	}
	
	response, err := internal.Connection(url, cache)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for _, result := range response.Results {
		fmt.Println("\t", result.Name)
	}

	conf.Next = response.Next
	conf.Previous = response.Previous
	return nil
}

func CommandMapb(conf *Config, cache *internal.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if conf.Previous != nil {
		url = *conf.Previous
	} else {
		fmt.Println("You have reached the beginning")
	}
	
	response, err := internal.Connection(url, cache)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for _, result := range response.Results {
		fmt.Println("\t", result.Name)
	}

	conf.Next = response.Next
	conf.Previous = response.Previous
	return nil
}
