package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	var id string = "12"

	personDB["12"] = PersonInfo{"12", "fff", "sss"}
	personDB["23"] = PersonInfo{"23", "dddd", "eeee"}

	delete(personDB, "12")

	person, ok := personDB[id]

	if ok {
		fmt.Println("Found person", person.Name)
	} else {
		fmt.Println("Dit not find person with id:" + id)
	}

	value, ok := personDB["23"]
	if ok {
		fmt.Print(value)
	}
}
