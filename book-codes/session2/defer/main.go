package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	f()
	defer func() {
		fmt.Println("defer with annonymous function")
	}()
}

func f() {

	defer fmt.Println("func f finished")
	fmt.Println("func f")
}

// output :

// func f
// func f finished
// defer with annonymous function
// 3
// 2
// 1
