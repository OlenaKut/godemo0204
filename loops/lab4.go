package loops

import (
	"fmt"
)

func Loops4() {

	var num1, sum int
	for i := 0; i <= 10; i++ {
		fmt.Printf("Type a number:")
		fmt.Scanln(&num1)

		sum += num1
	}

	fmt.Printf("%d\n", sum)
}
