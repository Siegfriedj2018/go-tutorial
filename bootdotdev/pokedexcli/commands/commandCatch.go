package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
)

func CommandCatch(conf *Config, cash *internal.Cache, pokemon ...string) error {
	url := internal.BaseURL + pokemon[0]

	fmt.Println("Throwing a Pokeball at %s...", pokemon[0])
	res, err := internal.CatchPokemon(url, cash)
	

	return nil
}