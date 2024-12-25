package main

import (
	"database/sql"
	"log"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// database handle--Dont do this in production
var db *sql.DB

type Album struct {
	ID		 int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties
	cfg := mysql.Config{
		User:		"root",
		Passwd: "",  // insert root password
		Net:		"tcp",
		Addr: 	"127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}