package main

import "fmt"

// Sum 1 to n  alghoritm

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

// Show Detail :
func ShowDetail(f func() (string, int, string)) {

	name, price, ingredients := f()

	fmt.Printf(" name: %s \n price: %d \n ingredients: %s ", name, price, ingredients)
}
func burger() (string, int, string) {
	return "Burger", 185000, "tomato + meat , cheese"
}

func main() {
	fmt.Println(sum(10)) // output : 55
	ShowDetail(burger)
}
