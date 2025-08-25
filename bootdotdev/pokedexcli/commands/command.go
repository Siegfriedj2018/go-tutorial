package commands

type cliCommand struct {
	Name				string
	Description string
	Callback		func(*Config) error
}

type Config struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


func GetCommands(conf *Config) map[string]cliCommand {
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
			Description: 	"Displays the next 20 locaions",
			Callback: 		CommandMap,	
		},
		"mapb": {
			Name:					"mapb",
			Description: 	"Displays the previous 20 locations",
			Callback: 		CommandMap,
		},
	}
}
