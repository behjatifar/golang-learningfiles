package main

import "fmt"

func main() {

	var ages = make(map[string]int)

	ages["parsa"] = 18
	ages["puran"] = 25
	ages["Azerakhsh"] = 10

	fmt.Println(ages)

	// for iterates over map first variable returns key & second for value
	for key, value := range ages {
		fmt.Println(key, value)
	}

}

//
