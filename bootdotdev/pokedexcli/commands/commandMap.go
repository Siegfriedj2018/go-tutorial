package commands

import (
	"fmt"


)

func CommandMap(conf *Config) error {
	if conf.Next == nil {
		return fmt.Errorf("you are on the last page, hint: try mapb")
	}

	return displayLocation(conf, *conf.Next)
}

func displayLocation(conf *Config, url string) error {

	return nil
}