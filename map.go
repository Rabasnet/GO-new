package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"california": 123456,
		"texas":      321654,
		"newyork":    987654,
		"louisiana":  456987,
	}
	fmt.Println(len(statePopulations))
	fmt.Println(statePopulations)
	statePopulations["georgia"] = 254689 //adding value to the map
	fmt.Println(len(statePopulations))
	delete(statePopulations, "texas") //deleting value of the map
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
}
