package main

import "fmt"

func main() {
	// strings contains lowercase and uppercase
	str := "abCde"

	// prepare the rune of strings
	runes := []rune(str)

	// iterates the strings
	for i, s := range str {
		// check each of character if it's lowercase
		// and convert to uppercase
		runes[i] = toUpper(s)
	}

	// print
	fmt.Println(string(runes))
}

// toUpper is a function to check the rune
// is lowercase then convert to uppercase
func toUpper(s rune) rune {
	// check if it's lowercase
	if s >= 'a' && s <= 'z' {
		// convert to uppercase
		s = s - 'a' + 'A'
	}

	return s
}
