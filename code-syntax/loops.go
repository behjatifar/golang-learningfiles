package main

import "fmt"

// Loop syntax in go
func main() {
	fruitsArray := []string{"apple", "banana", "cherry"}

	for _, fruits := range fruitsArray {
		fmt.Println(fruits)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	b := 0
	for b < 5 {
		b++
	}

	for _, item := range fruitsArray {
		fmt.Println(item)
	}

}
