package commands

type cliCommand struct {
	Name				string
	Description string
	Callback		func() error
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
	}
}
