package main

import (
	"fmt"
	"math"
)

func main() {
	var numBer int
	var total, pow float64

	fmt.Scanf("%d", &numBer)

	for i := 1; i <= numBer; i++ {
		fmt.Scanf("%v", &pow)
		intPow := math.Mod(pow, 10)
		pow = math.Floor(pow / 10)
		power := math.Pow(pow, intPow)
		total += power
	}
	fmt.Println(int(math.Round(total)))
}
