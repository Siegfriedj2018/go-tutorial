package main

import "fmt"

func main () {
	var publisher, writer, artist, title string
	var year, pageNumber uint16
	var grade float32

	title, writer = "Mr. GoToSleep", "Tracey Hatchet"
	artist, publisher, year, pageNumber, grade = "Jewel Tampson", 
																							 "DizzyBooks Publishing Inc.",
																								1997,
																								14,
																								6.5

	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println(publisher, "published in", year, "with", pageNumber, "pages")
	fmt.Println("the following grade was given to this book", grade)
	fmt.Println()
	title, writer = "Epic Vol. 1", "Ryan N. Shawn"
	artist, publisher, year, pageNumber, grade = "Phoebe Paperclips",
																							 "DizzyBooks Publishing Inc.",
																							 2013,
																							 160,
																							 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println(publisher, "published in", year, "with", pageNumber, "pages")
	fmt.Println("the following grade was given to this book", grade)
}