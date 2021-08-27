package stopwatch

import (
	"fmt"
)

func main() {
	var numButtonPresses int
	var time int
	var totaltime int

	fmt.Scanf("%d", &numButtonPresses)

	for s := 1; s <= numButtonPresses; s++ {
		fmt.Scanf("%d", &time)
		if s%2 != 0 {
			totaltime -= time
		} else {
			totaltime += time
		}

	}

	if numButtonPresses % 2 != 0 {
		fmt.Println("still running")
	} else {
		fmt.Println(totaltime)
	}
}
