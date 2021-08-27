package main

import (
	"fmt"
)

func main() {
	var daysNumber int

	fmt.Scanf("%d", &daysNumber)

	for s := 1; s <= daysNumber; s++ {
		fmt.Println(s, "Abracadabra")
	}
}
