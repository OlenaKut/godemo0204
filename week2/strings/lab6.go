package strings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func String6() {

	var text string

	fmt.Print("Ange text:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToLower(text)
	text2 := ""

	for i := len(text) - 1; i >= 0; i-- {
		text2 += string(text[i])
	}

	if text2 == text {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Ej palindrom")
	}

}

/*func String6() {
	var text string

	fmt.Print("Ange text:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToLower(text)

	// Use a rune slice for better performance
	runes := []rune(text)
	reversed := make([]rune, len(runes))

	// Reverse the string efficiently
	for i := len(runes) - 1; i >= 0; i-- {
		reversed[len(runes)-1-i] = runes[i]
	}

	// Compare original and reversed
	if string(reversed) == text {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Ej palindrom")
	}
}*/
