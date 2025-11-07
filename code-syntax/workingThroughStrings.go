package main

import (
	"fmt"
	"strings"
)

func main() {

	myString := "this is golang tutorial go go"

	fmt.Println("Contains", strings.Contains(myString, "go1"))
	fmt.Println("ContainAny", strings.ContainsAny(myString, "go1"))
	fmt.Println("Count", strings.Count(myString, "go"))
	fmt.Println(strings.Cut(myString, "go"))

	myStringArray := strings.Split(myString, " ")

	println("Word Count: ", len(myStringArray))

	for _, item := range myStringArray {
		fmt.Println(item)
	}

	myStringArray2 := strings.Join(myStringArray, ",")
	fmt.Println("Join ", myStringArray2)

	fmt.Println(strings.Repeat("persian gulf for ever", 10))

	println("replace", strings.Replace(myString, "go", "golang", 2))

	// fmt.Println(myStringArray)

}
