package main

import (
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/config"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to the blog aggregator.")
	fmt.Println("Reading config...")
	currentState := &config.State{
		Conf: config.Read(),
	}

	fmt.Printf("Read config: %+v\n", currentState.Conf)

	allCmds := config.Commands{}
	allCmds.Register("login", config.HandlerLogin)

	if len(os.Args) <= 1 {
		log.Fatalf("please provide a command. e.g. 'login <username>'")
	}
	cmdArgs := config.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	
	err := allCmds.Run(currentState, cmdArgs)
	if err != nil {
		log.Fatalf("there was an error running the command: %v", err)
	}
}
