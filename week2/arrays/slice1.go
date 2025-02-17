package arrays

import (
	"fmt"
	"strings"
)

func Slice1() {

	allNames := make([]string, 0, 10)

	var name string
	for {
		fmt.Printf("Type your name:")
		fmt.Scanln(&name)

		allNames = append(allNames, name)

		var answerInput string
		fmt.Print("Would you like to continue? (Y/N)")
		fmt.Scanln(&answerInput)
		if strings.ToLower(answerInput) != "y" {
			break
		}
	}

	longest := allNames[0]
	for _, val := range allNames {
		if len(val) > len(longest) {
			longest = val
		}
	}

	fmt.Printf("Longest is %s\n", longest)
}
