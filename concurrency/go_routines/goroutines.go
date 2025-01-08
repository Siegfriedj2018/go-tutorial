package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main () {
	// go keyword starts a goroutine in the same address space
	// this has to be synchronized for shared memory
	// go routine is a lightweight thread managed by Go
	go say("world")
	say("hello ")
}