package main

import (
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/config"
	"log"
)
func main() {
	fmt.Println("Welcome to the blog aggregator.")
	fmt.Println("Reading config...")
	var currentState config.State

	
	currentState.Conf = config.Read()
	fmt.Printf("Read config: %+v\n", conf)
	
	
	
	fmt.Println("Reading config again...")
	conf = config.Read()
	fmt.Printf("Read config: %+v\n", conf)
}