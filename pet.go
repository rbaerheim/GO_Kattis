package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3, num4, temptotal, total, winner int

	for i := 1; i <= 5; i++ {
		fmt.Scanf("%d %d %d %d", &num1, &num2, &num3, &num4)
		temptotal = num1 + num2 + num3 + num4
		if temptotal > total {
			total = temptotal
			winner = i
		} else {
			continue
		}

	}
	fmt.Println(winner, total)
}
