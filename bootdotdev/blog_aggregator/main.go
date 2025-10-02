package main

import (
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/config"
	"log"
	"os"
	_ "github.com/lib/pq"
)

type state struct {
	Conf  *config.Config
}


func main() {
	fmt.Println("Welcome to the blog aggregator.")
	fmt.Println("Reading config...")
	currentState := &state{
		Conf: config.Read(),
	}

	fmt.Printf("Read config: %+v\n", currentState.Conf)

	allCmds := commands{}
	allCmds.Register("login", handlerLogin)

	if len(os.Args) <= 1 {
		log.Fatalf("please provide a command. e.g. 'login <username>'")
	}
	cmdArgs := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err := allCmds.Run(currentState, cmdArgs)
	if err != nil {
		log.Fatalf("there was an error running the command: %v", err)
	}
}
