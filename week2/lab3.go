package week2

import (
	"fmt"
)

func Ifs3() {
	var num1, num2, bigger int

	fmt.Printf("Type any number: ")
	fmt.Scanln(&num1)

	fmt.Printf("Type one more number: ")
	fmt.Scanln(&num2)

	if num1 < num2 {
		bigger = num2
		fmt.Printf("Bigger number is %d\n", bigger)
	} else if num1 > num2 {
		bigger = num1
		fmt.Printf("Bigger number is %d\n", bigger)
	} else {
		fmt.Printf("Numbers are equal\n")
	}
}
