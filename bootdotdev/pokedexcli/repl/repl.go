package repl

import (
	"fmt"
	"bufio"
	"os"

	"go-tutorial/bootdotdev/pokedexcli/pokedex"
	"go-tutorial/bootdotdev/pokedexcli/commands"
)

func StartRepl(conf *commands.Config) {
	// nextUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	// conf := commands.Config{
	// 	Next: 		&nextUrl,
	// 	Previous: nil,
	// }

	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput, _ := pokedex.CleanInput(scanner.Text())
		
		cmd, ok := commands.GetCommands(conf)[cleanedInput[0]]
		if ok {
			err := cmd.Callback(conf)
			if err != nil {
				fmt.Printf("an unexpected error happened: %v\n", err)
				os.Exit(1)
			}
		}else {
			fmt.Printf("Unknown Command: %s\n", cleanedInput[0])
		}
	}
}

