package week1

import "fmt"

func Run5() {
	var name string

	fmt.Printf("Type your name: ")
	fmt.Scanln(&name)

	nameFive := name + name + name + name + name

	fmt.Printf("Your name is %s \n", nameFive)

}
