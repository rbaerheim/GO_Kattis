package FizzBuzz

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	for i := 1; i <= num3; i++ {
		if i%num1 == 0 && i%num2 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%num2 == 0 {
			fmt.Println("Buzz")
		} else if i%num1 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
