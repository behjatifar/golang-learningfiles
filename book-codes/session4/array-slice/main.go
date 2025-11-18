package main

import "fmt"

func main() {

	// how to define a array empty array

	studentsName := [5]string{}
	var studentAges [5]int

	fmt.Println(studentAges, studentsName) // expecting "" in strings and 0 in ints (zero values rule)

	studentNames := [5]int{10, 11, 15}

	for i := 0; i < 4; i++ {
		fmt.Print(studentNames[i])
	}

	// define slice

	numbers := make([]int, 8, 16)

	numbers = append(numbers, 10)

	// s := make([]int, 10, 15)
	// fmt.Println(cap(numbers))

	nums := []int{10, 11, 123, 131, 2, 12}

	nums = append(nums[:1], append([]int{60}, nums[1:]...)...)

	// with : you can accsess to specefic part of slice
	// point  the last index is not included. for example nums[0:3] would include this indexes  [0,1,2]
	fmt.Println("\n", nums[0:3]) // output : 10 , 11 ,123  index [0,,2]

	//  variadic functions are functions that in parameters they have slice
	// it means you are not limited to pass how many parameters as a type

	// e.g
	numbers2 := []int{1, 2, 3}
	fmt.Println(sum(numbers2...))
}
func sum(numbers ...int) int {

	var sum int

	for i := 0; i < len(numbers); i++ {
		sum += i
	}
	return sum
}
