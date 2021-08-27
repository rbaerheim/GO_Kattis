package sorttwonumbers

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int

	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	if num1 > num2 {
		fmt.Println(num2, num1)
	} else {
		fmt.Println(num1, num2)
	}
}
