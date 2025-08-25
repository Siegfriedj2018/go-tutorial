package main

import (
	"bootdevproject/pokedexcli/internal"
	"bootdevproject/pokedexcli/repl"
)


func main() {
	conf := internal.Connection()
	repl.StartRepl(conf)
}