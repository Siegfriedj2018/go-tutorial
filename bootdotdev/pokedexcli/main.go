package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/siegfriedj2018/bootdotdev/pokedexcli/pokedex"
)
var commands map[string]cliCommand

type cliCommand struct {
	name				string
	description string
	callback		func() error
}
func init() {
	commands = map[string]cliCommand {
		"exit": {
			name:					"exit",
			description:	"Exit the Pokedex",
			callback:			commandExit,
		},
		"help": {
			name:					"help",
			description: 	"Displays this help message",
			callback: 		commandHelp,	
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := pokedex.CleanInput(userInput)
		cmd, ok := commands[cleanedInput[0]]
		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("an unexpected error happened: %v\n", err)
				os.Exit(1)
			}
		}else {
			fmt.Printf("Unknown Command: %s\n", cleanedInput[0])
		}


	}
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standard input errored:", err)
	// }
}