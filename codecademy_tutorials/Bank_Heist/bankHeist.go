package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	rand.NewSource(time.Now().UnixNano())
	isHeistOn := true
	
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Past Guards...")
	} else {
		isHeistOn = false
		fmt.Println("Plan has failed, better luck next time.")
	}

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and Go!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened, Try again later...")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Where did the extra guards come from...")
		case 1:
			isHeistOn = false
			fmt.Println("The vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("who sounded the silent alarm...")
		case 3:
			isHeistOn = false
			fmt.Println("You had a change of heart and decided against robbing...")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 1000 + rand.Intn(1000000)
		fmt.Println("The total amount stolen was", amtStolen)
	}

	fmt.Println("The heist is currently:", isHeistOn)

}