package commands

import(
	"fmt"
	
)

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}