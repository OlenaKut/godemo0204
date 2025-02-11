package week2

import (
	"fmt"
	"strings"
)

func Ifs8() {

	var money int
	var discountiInput string
	var discount bool

	fmt.Printf("How much money do you have? ")
	fmt.Scanln(&money)

	fmt.Printf("Do you have a discount? ")
	fmt.Scanln(&discountiInput)

	if strings.ToLower(discountiInput) == "yes" || strings.ToLower(discountiInput) == "y" {
		discount = true

	} else if strings.ToLower(discountiInput) == "no" || strings.ToLower(discountiInput) == "n" || strings.ToLower(discountiInput) == " " {
		discount = false
	}

	if money >= 15 && money <= 25 && !discount {
		fmt.Printf("You can buy an hamburger ")
	} else if money >= 15 && money <= 25 && discount {
		fmt.Printf("You can buy an hamburger and a potato ")
	} else if money > 25 && money < 50 && !discount {
		fmt.Printf("You can buy a big hamburger ")
	} else if money > 25 && money < 50 && discount {
		fmt.Printf("You can buy a big hamburger and a potato")
	} else if money >= 50 && discount || money > 60 {
		fmt.Printf("You can buy everything ")
	}
}

//if money >= 15 && money <= 25 {
	//if discount {
		//fmt.Printf("You can buy an hamburger and a potato ")
	//} else {
		//fmt.Printf("You can buy an hamburger ")
	//}
//} else if money > 25 && money < 50 {
	//if discount {
		//fmt.Printf("You can buy a big hamburger and a potato")
	//} else {
		//fmt.Printf("You can buy a big hamburger ")
	//}
//} else if money money > 60 || (money >= 50 && money <= 60 && discount) {
	//fmt.Printf("You can buy everything ")
//}
