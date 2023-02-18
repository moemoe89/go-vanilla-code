package main

import "fmt"

func main() {
	// strings contains lowercase and uppercase
	str := "ABcDE"

	// prepare the rune of strings
	runes := []rune(str)

	// iterates the strings
	for i, s := range str {
		// check each of character if it's uppercase
		// and convert to lowercase
		runes[i] = toLower(s)
	}

	// print
	fmt.Println(string(runes))
}

// toLower is a function to check the rune
// is uppercase then convert to lowercase
func toLower(s rune) rune {
	// check if it's uppercase
	if s >= 'A' && s <= 'Z' {
		// convert to lowercase
		s = s - 'A' + 'a'
	}

	return s
}
