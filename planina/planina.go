package main

import (
	"fmt"
	"math"
)

func main() {
	var iterations int

	fmt.Scanf("%d", &iterations)

	finish := 2

	for i := 1; i <= iterations; i++ {
		finish = finish + finish - 1
	}

	sum := int(math.Pow(float64(finish), 2))

	fmt.Println(sum)
}
