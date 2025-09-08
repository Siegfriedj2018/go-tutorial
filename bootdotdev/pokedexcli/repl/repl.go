package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"go-tutorial/bootdotdev/pokedexcli/commands"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func StartRepl() {
	conf := &commands.Config{}
	pokeCash := *internal.NewCache(5 * time.Minute)
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := pokedex.CleanInput(scanner.Text())
		
		cmd, ok := commands.GetCommands()[cleanedInput.Current]
		if ok {
			err := cmd.Callback(conf, &pokeCash, cleanedInput.UserInput...)
			if err != nil {
				fmt.Printf("an unexpected error happened: %v\n", err)
				os.Exit(1)
			}
		}else {
			fmt.Printf("Unknown Command: %s\n", cleanedInput.Current)
		}
	}
}

