package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var number string
	var total, totalFromLineNums int

	fmt.Scanf("%s", &number)

	numbers := strings.Split(number, ";")

	for _, nums := range numbers {
		if strings.Contains(nums, "-") {
			numswithLine := strings.Split(nums, "-")
			for _, newNums := range numswithLine {
				intnums, err := strconv.Atoi(newNums)
				if intnums > 1000000 {
					fmt.Println(err)
				}
				if totalFromLineNums == 0 {
					totalFromLineNums -= intnums
				} else if totalFromLineNums < 0 {
					totalFromLineNums += intnums + 1
					total += totalFromLineNums
					totalFromLineNums = 0
				}
			}
		} else {
			intnums, err := strconv.Atoi(nums)
			if intnums > 100000 {
				fmt.Println(err)
			}
			total += 1
		}

	}
	fmt.Println(total)
}
