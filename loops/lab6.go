package loops

import (
	"fmt"
	"math/rand"
	"strings"
)

func Loops6() {
	const min = 1
	const max = 6

	for {
		num1 := rand.Intn(max-min+1) + min
		num2 := rand.Intn(max-min+1) + min

		fmt.Printf("%d och %d\n", num1, num2)

		var answerInput string
		fmt.Printf("Roll the dices again? (Y/N)")
		fmt.Scanln(&answerInput)
		if strings.ToLower(answerInput) != "y" {
			break
		}
	}
}
