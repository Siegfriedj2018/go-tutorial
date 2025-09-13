package pokedex

import (
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"log"
	"strings"
	"time"
)
type CurrentCommand struct {
	CMD				string
	ExtraCMD	[]string
}


// Add a way to store the date the pokemon was caught
type Pokedex struct {
	DateCreated			time.Time
	CaughtPokemon 	map[string]internal.PokemonDetails
}

func NewPokeDex() *Pokedex {
	log.Println("Creating Pokedex...")
	poke := &Pokedex{
		DateCreated: 		time.Now(),
		CaughtPokemon: 	make(map[string]internal.PokemonDetails),
	}

	return poke
}

func (poo *Pokedex) Add(key string, pooks internal.PokemonDetails) {
	log.Println("adding pokemon to database")
	poo.CaughtPokemon[key] = pooks
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

