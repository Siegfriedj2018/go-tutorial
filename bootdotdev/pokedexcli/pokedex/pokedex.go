package pokedex

import (
	"strings"
)

func CleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lowChar := strings.ToLower(text)
	output := strings.Fields(lowChar)
	return output
}