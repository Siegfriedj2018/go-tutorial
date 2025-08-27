package repl

import (
	"bufio"
	"fmt"
	"os"

	"go-tutorial/bootdotdev/pokedexcli/commands"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func StartRepl() {
	conf := &commands.Config{}
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

