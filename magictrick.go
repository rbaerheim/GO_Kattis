package main

import (
	"fmt"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var input string
	var letterlist []string
	var finished int

	fmt.Scanf("%s", &input)

	for _, letters := range input {
		letter := string(letters)
		if contains(letterlist, letter) {
			fmt.Println("0")
			finished = 1
			break
		}
		letterlist = append(letterlist, letter)
	}
	if finished != 1 {
		fmt.Println("1")
	}
}
