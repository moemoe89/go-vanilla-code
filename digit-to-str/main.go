package main

import "fmt"

func main() {
	// only works for single digit 0-9
	digit := 9

	// need convert int to rune
	// '0' + -> to convert as rune string
	// then convert to string
	str := string('0' + rune(digit))

	// print
	fmt.Println(str)
}
