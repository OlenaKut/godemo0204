package week2

import "fmt"

func Ifs6() {

	var age int

	fmt.Printf("Type your age: ")
	fmt.Scanln(&age)

	if age >= 1980 && age < 1990 {
		fmt.Printf("You was born in the 1980s\n")
	} else if age >= 1990 && age < 2000 {
		fmt.Printf("You was born in the 1990s\n")
	} else {
		fmt.Printf("You wasn't born in the 1980s or 1990s\n")

	}

}
