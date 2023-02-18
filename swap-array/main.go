package main

import "fmt"

func main() {
	// prepare the original array
	swap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// prepare the first and last index
	i := 0
	j := len(swap) - 1

	// iterates until i > j
	for i < j {
		// swap the first with the last index
		// and moving one the next index (left to right, right to left)
		swap[i], swap[j] = swap[j], swap[i]

		// increment and decrement
		i++
		j--
	}

	// print the swap array
	fmt.Println(swap)
}
