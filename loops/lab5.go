package loops

import (
	"fmt"
)

func Loops5() {

	var num1 int
	for i := 0; i <= num1; i++ {
		fmt.Printf("Type a number:")
		fmt.Scanln(&num1)

	}

	fmt.Printf("%d\n", num1)
}
