package commands

import (
	"fmt"
)

func CommandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands(conf) {
		fmt.Printf("\t%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}