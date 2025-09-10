package pokedex

import (
	"strings"
)
type CurrentCommand struct {
	CMD				string
	ExtraCMD	[]string
}

func CleanInput(text string) CurrentCommand {
	if len(text) == 0 {
		return CurrentCommand{}
	}
	lowChar := strings.ToLower(text)
	output := strings.Fields(lowChar)
	curCMD := CurrentCommand{
		CMD:			output[0],
		ExtraCMD:	output[1:],
	}
	return curCMD
}
