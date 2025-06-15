package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("You have this much fuel left in Gallons:", fuel)
}

func calculatefuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
} 

func greetPlanet(planet string) {
	fmt.Println("Hello and Welcome to", planet)
}

func cantFly() {
	fmt.Println("WE do not have the avaiable fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	 var fuelRemaining int
	 var fuelCost int
	 fuelRemaining = fuel

	 fuelCost = calculatefuel(planet)
	 if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	 }

	 if fuelCost > fuelRemaining {
		cantFly()
	 }

	 return fuelRemaining
} 

func main() {
	var fuel int
	var planetChoice string
	fuel = 1000000

	planetChoice = "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}