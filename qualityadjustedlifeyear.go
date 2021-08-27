package main

import (
	"fmt"
)

func main() {
	var qualityNum int

	var lifeQuality, yearNum, accuMulated float64

	fmt.Scanf("%d", &qualityNum)

	for i := 1; i <= qualityNum; i++ {
		fmt.Scanf("%v %v", &lifeQuality, &yearNum)
		accuMulated += lifeQuality * yearNum
	}

	fmt.Println(accuMulated)
}
