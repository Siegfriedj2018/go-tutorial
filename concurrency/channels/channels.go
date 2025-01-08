package main

import "fmt"

func sum (s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}



func main() {
	// Buffered channels example
	// ch := make(chan int, 3)
	// ch <- 1
	// ch <- 2
	// ch <- 3

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)  // creates the channel
	go sum(s[:len(s)/2], c) // sums the first half
	go sum(s[len(s)/2:], c) // sums the second half
	x, y := <-c, <-c // receive from c (like join?)
	fmt.Println(x, y, x+y)
}