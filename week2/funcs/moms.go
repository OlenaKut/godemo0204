package funcs

import "fmt"

func Moms() {

	var num, vat float64

	fmt.Printf("Enter the price: ")
	fmt.Scanln(&num)

	fmt.Printf("Enter the VAT rate: ")
	fmt.Scanln(&vat)

	result := CountVat(num, vat)
	fmt.Println(result)
}

func CountVat(f1, f2 float64) float64 {
	return (f1 * f2) / 100
}
