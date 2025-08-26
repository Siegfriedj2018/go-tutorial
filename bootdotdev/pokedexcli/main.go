package main

import (
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/repl"
)


func main() {
	conf := internal.Connection()
	repl.StartRepl(conf)
}