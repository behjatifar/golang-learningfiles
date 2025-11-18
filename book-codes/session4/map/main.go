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

	// delete a value
	delete(ages, "parsa")

	// check specefic key is valid or not
	age, ok := ages["maryam"]
	age, ok = ages["sadegh"]
	age, ok = ages["Azerakhsh"]

	if ok != true {
		fmt.Println("user not valid")
	} else {
		fmt.Println(age)
	}

}

//
