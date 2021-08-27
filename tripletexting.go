package tripletexting

import (
	"fmt"
	_ "strings"
)

func main() {
	var input string
	// var wordList []string

	fmt.Scanf("%s", &input)

	lenghtAll := len(input)

	lenghtOne := lenghtAll / 3

	firstWord := input[:lenghtOne]
	secondWord := input[lenghtOne:(lenghtOne + lenghtOne)]
	thirdWord := input[(lenghtOne + lenghtOne):(lenghtOne + lenghtOne + lenghtOne)]

	if firstWord == secondWord {
		fmt.Println(firstWord)
	} else if secondWord == thirdWord {
		fmt.Println(secondWord)
	} else if firstWord == thirdWord {
		fmt.Println(thirdWord)
	} else {
		fmt.Println(firstWord)
	}
}
