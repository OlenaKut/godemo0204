package funcs

import "fmt"

func Plus() {
	var text1, text2 string

	fmt.Printf("Type a text: ")
	fmt.Scanln(&text1)

	fmt.Printf("Type one more text: ")
	fmt.Scanln(&text2)

	//Option 1
	//text := text1 + " " + text2
	//fmt.Println(text)

	//Option 3
	//result := fmt.Sprintf("%s, %s\n", text1, text2)
	//fmt.Println(result)

	//Option 3
	result := ConcatenateStrings(text1, text2)
	fmt.Println(result)

}

func ConcatenateStrings(s1, s2 string) string {
	return s1 + " " + s2
}
