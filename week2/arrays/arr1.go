package arrays

import "fmt"

func Arr1() {
	var num1 [4]int

	for index := 0; index < 4; index++ {
		fmt.Print("Type a number: ")
		fmt.Scanln(&num1[index])
	}

	bigger := num1[0]
	for index := 1; index < 4; index++ {
		if num1[index] > bigger {
			bigger = num1[index]
		}
	}

	fmt.Printf("Bigger is %d\n", bigger)
}
