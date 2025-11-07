package main

import "fmt"

func main() {
	tahere := 0
	ebrahim := 0

	tahereScore := func() int {
		tahere++
		return tahere
	}
	ebrahimScore := func() int {
		ebrahim++
		return ebrahim
	}

	fmt.Println(tahereScore())
	fmt.Println(tahereScore())
	fmt.Println(ebrahimScore())
}
