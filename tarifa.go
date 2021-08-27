package tarifa

import (
	"fmt"
)

func main() {
	var megabyteseachMonth int
	var numberofMonths int
	var usedMBthismonth int

	fmt.Scanf("%d", &megabyteseachMonth)
	fmt.Scanf("%d", &numberofMonths)

	i := 0

	totalMBleft := megabyteseachMonth

	for i != numberofMonths {
		totalMBleft += megabyteseachMonth
		
		fmt.Scanf("%d", &usedMBthismonth)

		if totalMBleft < usedMBthismonth {
			break
		}

		totalMBleft -= usedMBthismonth

		i = i + 1
	}

	fmt.Println(totalMBleft)
}
