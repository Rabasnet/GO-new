package main

import "fmt"

//struct can encompass multiple data types unlike others
//Doctor is visible in other packages but the variables under Doctor struct are not visible
//To make it visible, the variables also needs to be capitilized
type Doctor struct {
	number     int
	actorName  string `required max:"100"` //using tags
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "john wick",
		companions: []string{
			"rabi",
			"babi",
			"jo",
		},
	}
	fmt.Println(aDoctor.companions[2])
}
