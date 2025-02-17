package strings

import (
	"fmt"
	"strings"
)

func CapitalizeEachWord(input string) string {
	words := strings.Fields(input) // Split by spaces
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func String3() {
	var text = "kurt andersson"
	fmt.Println(CapitalizeEachWord(text)) // Output: Kurt Andersson
}
