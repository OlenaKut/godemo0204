package strings

import (
	"fmt"
	"strings"
)

func String4() {
	var text = "This is a text I need to change"

	text = strings.ReplaceAll(text, " ", "*")

	total := strings.Count(text, "*")
	fmt.Printf("%s, %d\n", text, total)
}
