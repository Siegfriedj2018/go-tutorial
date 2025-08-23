package repl

import (
	"fmt"
	"bufio"
	"os"

	"bootdevproject/pokedexcli/pokedex"
	"bootdevproject/pokedexcli/commands"
)

func StartRepl() {
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := pokedex.CleanInput(userInput)
		cmd, ok := commands.GetCommands()[cleanedInput[0]]
		if ok {
			err := cmd.Callback()
			if err != nil {
				fmt.Printf("an unexpected error happened: %v\n", err)
				os.Exit(1)
			}
		}else {
			fmt.Printf("Unknown Command: %s\n", cleanedInput[0])
		}


	}
}

