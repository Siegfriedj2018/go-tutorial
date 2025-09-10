package commands

import (
	"fmt"
	
	"go-tutorial/bootdotdev/pokedexcli/internal"
)

func CommandExplore(conf *Config, cash *internal.Cache, locations ...string) error {
	if locations[0] == "" {
		return fmt.Errorf("You have not provided a location, type help for usage")
	}

	url := internal.BaseURL + "location-area/" + locations[0]

	fmt.Printf("Exploring %s...\n", locations[0])
	res, err := internal.RetrievePokemon(url, cash)
	if err != nil {
		return fmt.Errorf("there was an internal error: %w", err)
	}

	
	fmt.Println("Found Pokemon:")
	for idx, result := range res.PokemonEncounters {
		fmt.Printf("%d. %s\n", idx+1, result.Pokemon)
	}
	
	return nil
}