package pokedex

import (
	"strings"
)
type CurrentCommand struct {
	Current		string
	Previous	string
}

func CleanInput(text string) ([]string, CurrentCommand) {
	if len(text) == 0 {
		return []string{}, CurrentCommand{}
	}
	lowChar := strings.ToLower(text)
	output := strings.Fields(lowChar)
	curCMD := CurrentCommand{output[0], ""}
	return output, curCMD
}