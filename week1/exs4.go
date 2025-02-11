package week1

import "fmt"

func Run4() {
	var name string
	var surname string

	fmt.Printf("Type your name: ")
	fmt.Scanln(&name)

	fmt.Printf("Type your surname: ")
	fmt.Scanln(&surname)

	fmt.Printf("Your name is %s %s\n", name, surname)

}
