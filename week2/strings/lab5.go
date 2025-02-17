package strings

import (
	"fmt"
	"strings"
)

func String5() {

	var email string
	fmt.Print("Type your email: ")
	fmt.Scanln(&email)

	atIndex := strings.Index(email, "@")

	lastPoint := strings.LastIndex(email, ".")

	charsAfter := len(email) - lastPoint - 1

	if atIndex > 1 && (charsAfter == 2 || charsAfter == 3) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

}
