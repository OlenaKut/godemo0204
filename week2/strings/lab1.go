package strings

import "fmt"

func String1() {

	var text, sum string

	for i := 0; i < 3; i++ {
		fmt.Printf("Type any text: ")
		fmt.Scanln(&text)
		sum += text + " "
	}

	fmt.Printf("%s\n", sum)

}
