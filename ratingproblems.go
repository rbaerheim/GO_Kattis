package ratingproblems

import (
	"fmt"
)

func main() {
	var totalJudges, allreadyRated, ratingValues, totalPoints, judgesNotrated int

	fmt.Scanf("%d %d", &totalJudges, &allreadyRated)

	judgesNotrated = totalJudges - allreadyRated

	for i := 1; i <= allreadyRated; i++ {
		fmt.Scanf("%d", &ratingValues)
		totalPoints += ratingValues
	}

	minmax := judgesNotrated * 3

	floatmin := float64(minmax)
	floatmax := float64(minmax)
	floattotal := float64(totalPoints)
	floatjudges := float64(totalJudges)

	f64min := (floattotal - floatmin) / floatjudges
	f64max := (floattotal + floatmax) / floatjudges

	fmt.Printf("%v %v", f64min, f64max)
}
