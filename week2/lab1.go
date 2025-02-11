package week2

import (
	"fmt"
)

func Ifs1() {
	var num int
	fmt.Printf("Type a number:")
	fmt.Scanln(&num)

	if num >= 10 {
		fmt.Printf("Number is bigger then 10\n")
	} else if num < 10 {
		fmt.Printf("Number is less then 10\n")
	}

}
