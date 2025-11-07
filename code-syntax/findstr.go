package main

import (
	"fmt"
	"strings"
)

func main() {
	mystring := "hello world i'm here to adventure and enjoy"

	extractedText, startindex, endindex := findstr(mystring, "enjoy")

	if extractedText != nil && startindex != nil && endindex != nil {
		fmt.Println(*extractedText, *startindex, *endindex)
	} else {
		fmt.Println(*extractedText)
	}

}

func findstr(text string, search string) (*string, *int, *int) {
	start := strings.Index(text, search)
	if start == -1 {
		value := "not found text"
		return &value, nil, nil
	}
	end := start + len(search)
	value := text[start:end]

	return &value, &start, &end
}
