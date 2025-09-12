package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
)

func CommandCatch(conf *Config, cash *internal.Cache, pokemon ...string) error {
	url := internal.BaseURL + "pokemon/" + pokemon[0]

	fmt.Println("Throwing a Pokeball at %s...", pokemon[0])
	res, err := internal.CatchPokemon(url, cash)
	if err != nil {
		return fmt.Errorf("there was an internal error: %w", err)
	}

	catchChance := clamp(res.BaseExperience)
	




	return nil
}

func clamp(xp int) float64 {
	ceiling := 0.9
	floor := 0.1
	base_xp_mid := 300
	k := (ceiling - 0.5) / float64(base_xp_mid)
	
}