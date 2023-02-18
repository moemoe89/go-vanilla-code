package main

import "fmt"

func main() {
	// only positive numbers without leading zero
	// 012345, -12345 -> not acceptable
	n := 12345

	// prepare the array of numbers
	var numbers []int

	// iterates until numbers become 0
	for n != 0 {
		// use modulo 10 in order to get every digit
		// e.g. 12345 % 10 -> 5
		m := n % 10

		// use prepend because we need to add
		// digit from right to left
		// [5] -> [4 5] -> [3 4 5]
		numbers = append([]int{m}, numbers...)

		// divide the number to eliminate the last digit
		// 12345 to 1234
		n /= 10
	}

	// print
	fmt.Println(numbers)
}
