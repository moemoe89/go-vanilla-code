package main

import "fmt"

func main() {
	// prepare the array
	nums := []int{1, 6, 7, 3, 2, 1, 2, 45, 6, 1, 5, 45, 23, 1}

	// prepare the max frequent variable
	max := 0

	// prepare the map number and the frequent number
	m := make(map[int]int)

	// iterates the array
	for _, n := range nums {
		// set to map and count the frequent
		m[n]++

		// if the value more than the current max
		// replace the max value
		if m[n] > max {
			max = m[n]
		}
	}

	// array for the most frequent number
	var mostNums []int

	// iterates the map
	for k, v := range m {
		// check the value should be equal
		// with the amx number
		if v == max {
			// append to array
			mostNums = append(mostNums, k)
		}
	}

	// print
	fmt.Println(mostNums)
}
