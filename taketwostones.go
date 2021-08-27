package taketwostones

import (
	"fmt"
)

func main() {

	var stones int

	fmt.Scanf("%d", &stones)

	if stones%2 == 0 {
		fmt.Println("Bob")
	} else {
		fmt.Println("Alice")
	}
}
