package week2

import (
	"fmt"
)

func Ifs2() {
	var num int
	fmt.Printf("How many milk do you have:")
	fmt.Scanln(&num)

	if num < 10 {
		fmt.Printf("You need to order 30 packages\n")
	} else if num >= 10 && num <= 20 {
		fmt.Printf("You need to order 20 packages\n")
	} else {
		fmt.Printf("You don't need to order any milk\n")
	}

}
