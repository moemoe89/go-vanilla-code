package main

import "fmt"

func main() {
	// prepare the int array
	nums := []int{6, 9, 1, 2, 13, 66, 3, 4, 9}

	// count the index
	n := len(nums) - 1

	// iterates the array
	for i := 0; i <= n; i++ {
		// for optimized if the array already sorted
		isSwap := false

		// the GIF explanation for this algorithm can
		// be found on this link https://commons.wikimedia.org/wiki/File:Bubble-sort-example-300px.gif
		// Basically, bubble-sort will compare the current and next index value
		// and swap if the current more than next index value
		for j := 0; j < n-i; j++ {
			if nums[j] > nums[j+1] {
				isSwap = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		// break if not swapped
		if !isSwap {
			break
		}
	}

	// print
	fmt.Println(nums)
}
