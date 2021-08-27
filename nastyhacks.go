package nastyhacks

import (
	"fmt"
)

func main() {
	var cases, doNOTadvertise, doAdvertise, cost int

	fmt.Scanf("%d", &cases)

	for i := 1; i <= cases; i++ {
		fmt.Scanf("%d %d %d", &doNOTadvertise, &doAdvertise, &cost)
		if (doAdvertise - cost) > doNOTadvertise {
			fmt.Println("advertise")
		} else if (doAdvertise - cost) < doNOTadvertise {
			fmt.Println("do not advertise")
		} else {
			fmt.Println("does not matter")
		}
	}
}
