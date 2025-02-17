package funcs

import (
	"errors"
	"fmt"
)

func Tax() {
	var num float64
	fmt.Printf("Enter your salary:")
	fmt.Scanln(&num)

	tax, err := CalculateTaxesOnSalary(num)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Your tax is %f\n", tax)
	}

}

func CalculateTaxesOnSalary(num float64) (tax float64, err error) {
	if num < 0 {
		return 0, errors.New("inte negativ lön väl?")
	}
	if num < 15000 {
		return num * 0.12, nil
	}
	if num < 30000 {
		return num * 0.28, nil
	}
	return num * 0.33, nil
}
