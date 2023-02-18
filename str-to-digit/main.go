package main

import "fmt"

func main() {
	// only works for single digit string 0-9
	str := "0"

	// need convert string to rune using str[0]
	// - '0' > to convert as rune digit
	// then convert to int
	digit := int(str[0] - '0')

	// print
	fmt.Println(digit)
}
