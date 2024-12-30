package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

// database handle--Dont do this in production
var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist queries for albums that have the specified artist name.
// returns an album slice of artist that match the name
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	// using query to execute the select statement to query albums
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// Defer keeps rows so that any resources it holds will be released when the functions exits.
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}
	return albums, nil
}

func allAlbums() ([]Album, error) {
	var albums []Album

	row, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("allAlbums: %v", err)
	}

	defer row.Close()
	for row.Next() {
		var alb Album
		if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("allAlbum: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("allAlbums: %v", err)
	}
	return albums, nil

}

// albumById queries for the album with the specified id.
func albumByID(id int64) (Album, error) {
	// an album to hold data form the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album id of the new entry

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	// Capture connection properties
	cfg := mysql.Config{
		User:   "root",
		Passwd: "", // insert root password
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
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
	allAlbs, err2 := allAlbums()
	albums, err := albumsByArtist("John Coltrane")
	alb, err3 := albumByID(2)

	for _, alb := range allAlbs {
		fmt.Println(alb)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("John Coltrane albums found: %v\n", albums)

	// Hard-coded id 2 here to test the query
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "Endless Behavior",
		Artist: "Starset",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	allAlbs2, err4 := allAlbums()
	for _, albs := range allAlbs2 {
		fmt.Println(albs)
	}
	if err2 != nil {
		log.Fatal(err4)
	}

}
