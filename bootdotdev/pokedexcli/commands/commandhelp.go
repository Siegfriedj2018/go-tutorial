package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
)

func CommandHelp(conf *Config, cache *internal.Cache, extra ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("\t%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}