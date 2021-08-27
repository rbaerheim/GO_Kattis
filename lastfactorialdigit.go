package lastfactorialdigit

import (
	"fmt"
)

func main() {
	var inputNum int
	var factorial int

	fact := 1

	fmt.Scanf("%d", &inputNum)

	for i := 1; i <= inputNum; i++ {
		fmt.Scanf("%d", &factorial)
		if factorial == 1 {
			fact = 1
			fmt.Println(fact)
		} else {
			for j := 1; j <= factorial; j++ {
				fact *= j
			}
			facti := fmt.Sprint(fact)
			lastdigit := facti[len(facti)-1:]
			fmt.Println(lastdigit)
			fact = 1
		}
	}
}
