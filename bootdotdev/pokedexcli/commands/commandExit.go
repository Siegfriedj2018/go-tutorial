package commands

import (
	"os"
	"fmt"
)

func CommandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
