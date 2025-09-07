package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"os"
)

func CommandExit(conf *Config, cache *internal.Cache, extra ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
