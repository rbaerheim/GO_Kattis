package main

import (
	"fmt"
)

func main() {
	var lines, columns, classes int
	var student string
	var studNums []string

	fmt.Scanf("%d %d %d", &lines, &columns, &classes)

	for i := 1; i <= lines; i++ {
		fmt.Scanf("%s", &student)
		studNums = append(studNums, student)
	}
}
