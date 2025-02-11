package loops

import (
	"fmt"
	"strings"
)

func Loops3() {

	var num1, num2 int
	for {
		fmt.Printf("Type a number 1:")
		fmt.Scanln(&num1)

		fmt.Printf("Type a number 2:")
		fmt.Scanln(&num2)

		fmt.Printf("%d+%d=%d\n", num1, num2, num1+num2)

		var answerInput string
		fmt.Print("Would you like to continue? (Y/N)")
		fmt.Scanln(&answerInput)
		if strings.ToLower(answerInput) != "y" {
			break
		}
	}
}
