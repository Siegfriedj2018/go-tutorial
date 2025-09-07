package pokedex

import (
	"strings"
)
type CurrentCommand struct {
	Current		string
	UserInput	[]string
}

func CleanInput(text string) CurrentCommand {
	if len(text) == 0 {
		return CurrentCommand{}
	}
	lowChar := strings.ToLower(text)
	output := strings.Fields(lowChar)
	curCMD := CurrentCommand{
		Current:		output[0],
		UserInput:	output[1:],
	}
	return curCMD
}
