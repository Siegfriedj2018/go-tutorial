package main

import (
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/config"
	"log"
)
func main() {
	fmt.Println("Welcome to the blog aggregator.")
	fmt.Println("Reading config...")
	conf := config.Read()
	fmt.Printf("Read config: %+v\n", conf)
	
	err := conf.SetUser("justin")
	if err != nil {
		log.Fatalf("could not set current user: %v", err)
	}

	fmt.Println("Reading config again...")
	conf = config.Read()
	fmt.Printf("Read config: %+v\n", conf)
}