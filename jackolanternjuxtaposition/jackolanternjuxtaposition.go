package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3, answer int

	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	answer = num1 * num2 * num3
	fmt.Println(answer)
}
