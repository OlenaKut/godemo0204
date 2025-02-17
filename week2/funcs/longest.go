package funcs

import (
	"fmt"
	"strings"
)

func Longest() {
	allWords := make([]string, 0, 10)
	var text string

	for {
		fmt.Printf("Type a word:")
		fmt.Scanln(&text)

		allWords = append(allWords, text)

		var answerInput string
		fmt.Print("Would you like to continue? (Y/N)")
		fmt.Scanln(&answerInput)
		if strings.ToLower(answerInput) != "y" {
			break
		}
	}

	longest := FindTheLongest(allWords)
	fmt.Println(longest)
}

func FindTheLongest(allWords []string) (longest string) {
	longest = allWords[0]
	for _, val := range allWords {
		if len(val) > len(longest) {
			longest = val
		}
	}

	return longest
}
