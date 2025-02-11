package week2

import (
	"fmt"
	"strings"
)

func Ifs7() {
	var country string

	fmt.Printf("Where do you live: ")
	fmt.Scanln(&country)

	country = strings.ToLower(country)

	if country == "sweden" || country == "danmark" || country == "norwey" {
		fmt.Printf("You live in Scandinavia")
	} else {
		fmt.Printf("You don't live in Scandinavia")
	}
}
