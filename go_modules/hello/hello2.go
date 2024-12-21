package main

import (
	"fmt"
	"log"
	"example/greetings"
)

func main() {
	// Set properties of hte predefined Logger, including
	// the log entry prefix and a flag to disable printing
	//the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gemini", "Gippity", "Devon"}

	// Request a greeting message.
	// Updated: this line was updated for the addition of slices of names
	messages, err := greetings.Hellos(names)
	
	// If an error was returned, print it to the console
	// and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of 
	// messages to the console.
	for _, name := range messages {
		fmt.Println(name)
	}
}