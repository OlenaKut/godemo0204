package strings

import (
	"fmt"
	"strings"
)

func String2() {

	var text = "Hello, world"

	pos1 := strings.Index(text, "w")
	pos2 := strings.LastIndex(text, "o")

	fmt.Printf("Positions: %d and %d\n", pos1, pos2)

}
