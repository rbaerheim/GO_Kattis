package solvingforcarrots

import (
	"fmt"
)

func main() {
	var contestNum, carrotNum int
	var description string

	fmt.Scanf("%d %d", &contestNum, &carrotNum)

	for s := 1; s <= contestNum; s++ {
		fmt.Scanf("%s", description)
	}

	fmt.Println(carrotNum)
}
