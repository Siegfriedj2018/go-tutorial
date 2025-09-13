package commands

import (
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

type cliCommand struct {
	Name				string
	Description string
	Callback		func(*Config, *internal.Cache, *pokedex.Pokedex, ...string) error
}

type Config struct {
	Next			*string
	Previous 	*string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:					"exit",
			Description:	"Exit the Pokedex",
			Callback:			CommandExit,
		},
		"help": {
			Name:					"help",
			Description: 	"Displays this help message",
			Callback: 		CommandHelp,	
		},
		"map": {
			Name:					"map",
			Description: 	"Displays the next 20 locaions, if applicable",
			Callback: 		CommandMap,	
		},
		"mapb": {
			Name:					"mapb",
			Description: 	"Displays the previous 20 locations, if applicable",
			Callback: 		CommandMapb,
		},
		"explore": {
			Name:					"explore <location>",
			Description: 	"Displays the pokemon found in specified location",
			Callback: 		CommandExplore,
		},
		"catch": {
			Name:					"catch <Pokemon>",
			Description:  "Catches a Pokemon and adds it to the Pokedex",
			Callback:     CommandCatch,
		},
	}
}
