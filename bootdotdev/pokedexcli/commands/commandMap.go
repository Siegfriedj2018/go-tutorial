package commands

import (
	"fmt"
	"github.com/siegfriedj2018/bootdotdev/pokedexcli/internal"
)

func CommandMap(conf *Config) error {
	if conf.Next == nil {
		return fmt.Errorf("you are on the last page, hint: try mapb")
	}

	return internal.DisplayLocation(conf, conf.Next)
}