package main

import "fmt"

//  package variable that exists during of app running
var b = 5

func main() {
	// here we can't access to i from some function and if we try we would get errors
	// fmt.Println("we have i %d" i)

	//  so how we can access to a variable from function? exactly with pointers
	j := accessToVar()
	fmt.Println(*j)
}

func some() {
	// here we have i which is a  blocked scope variable
	var i int = 10
	// about life time of i , till some and this scope exist i is alive but after some executed it will be destroy from memory
	// i is intilized in stack
	fmt.Println(i)

}

func accessToVar() *int {
	x := 55

	return &x
}
