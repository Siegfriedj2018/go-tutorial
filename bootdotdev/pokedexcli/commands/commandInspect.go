package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func CommandInspect(conf *Config, cash *internal.Cache, pd *pokedex.Pokedex, pokeMon ...string) error {
	if len(pokeMon) == 0 {
		return fmt.Errorf("please enter a pokemon name, type help for usage.")
	}

	data, ok := pd.Get(pokeMon[0])
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	info := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", data.Name)
}