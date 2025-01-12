package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the server and set up the routes

	// serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting the server at http://localhost:4001")

	// start the server
}