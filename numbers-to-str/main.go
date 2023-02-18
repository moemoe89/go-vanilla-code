package main

import "fmt"

func main() {
	// only positive numbers without leading zero
	// 012345, -12345 -> not acceptable
	numbers := 12345

	// prepare the array of rune
	str := []rune{}

	// iterates until numbers become 0
	for numbers != 0 {
		// use modulo 10 in order to get every digit
		// e.g. 12345 % 10 -> 5
		m := numbers % 10

		// convert to rune string
		r := '0' + rune(m)

		// use prepend because we need to add
		// digit from right to left
		// [5] -> [4 5] -> [3 4 5]
		str = append([]rune{r}, str...)

		// divide the number to eliminate the last digit
		// 12345 to 1234
		numbers /= 10
	}

	// print
	fmt.Println(string(str))
}
