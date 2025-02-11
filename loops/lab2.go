package loops

import "fmt"

func Loops2() {

	var num1, num2 int
	fmt.Printf("Type a number 1:")
	fmt.Scanln(&num1)

	fmt.Printf("Type a number 2:")
	fmt.Scanln(&num2)

	for i := num1; i <= num2; i++ {
		fmt.Print(i)
	}
}
