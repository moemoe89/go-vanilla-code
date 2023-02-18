package main

import "fmt"

func main() {
	// only positive string numbers without leading zero
	// "012345" -> will become 12345
	// "-12345" -> not acceptable
	str := "-12345"

	// prepare the numbers variable
	numbers := 0

	// iterates string as rune
	for _, s := range str {
		// convert rune to integer by - '0'
		digit := int(s - '0')
		// multiply each digit with 10 to put
		// as the position and add the digit itself
		// e.g. 5 -> 5, 4 -> 40, 3 -> 300
		numbers = numbers*10 + digit
	}

	// print
	fmt.Println(numbers)
}
