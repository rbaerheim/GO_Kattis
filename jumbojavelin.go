package jumbojavelin

import (
	"fmt"
)

func main() {
	var numJavs, lenghtJavs, totalLenght int

	fmt.Scanf("%d", &numJavs)

	for i := 1; i <= numJavs; i++ {
		fmt.Scanf("%d", &lenghtJavs)
		if i == 1 {
			totalLenght = lenghtJavs
			continue
		}
		totalLenght += lenghtJavs - 1
	}
	fmt.Println(totalLenght)
}
