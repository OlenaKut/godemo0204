package week2

import (
	"fmt"
)

func Ifs4() {
	var num1, num2, num3, bigger int

	fmt.Printf("Type any number: ")
	fmt.Scanln(&num1)

	fmt.Printf("Type one more number: ")
	fmt.Scanln(&num2)
	fmt.Printf("And one more: ")
	fmt.Scanln(&num3)

	if num1 > num2 && num1 > num3{
		bigger = num1
		fmt.Printf("Bigger number is %d\n", bigger)
	} else if num2 > num1 && num2 > num3 {
		bigger = num2
		fmt.Printf("Bigger number is %d\n", bigger)
	} else if num3 > num1 && num3 > num2 {
		bigger = num3
		fmt.Printf("Bigger number is %d\n", bigger)
	} else {
		fmt.Printf("Numbers are equal\n")
	}
}