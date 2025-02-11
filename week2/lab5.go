package week2

import (
	"fmt"
	"strings"
)

func Ifs5() {
	var category string
	var price int

	fmt.Printf("Who are you (adult, student, pens)? ")
	fmt.Scanln(&category)

	category = strings.ToLower(category)

	if category == "adult" || category == "a" {
		price = 30
	} else if category == "student" || category == "s" || category == "pens" || category == "p" {
		price = 20
	}


	if price > 0 {
		fmt.Printf("You price is %d\n", price)
	} else {
		fmt.Println("You typed wrong category")
	}

}


