package main

import "fmt"

func main() {
	// sort int
	// prepare the int array
	nums := []int{6, 9, 1, 2, 13, 66, 3, 4, 9}

	// count the index
	n := len(nums) - 1

	// sort
	sortArr[int](n, nums)

	// print
	fmt.Println(nums)

	// sort string
	// prepare the int array
	strs := []string{"vcfgdf", "asdasda", "123asfdf", "xtyrhrt", "dtfgdfgh"}

	// count the index
	n = len(strs) - 1

	// sort
	sortArr[string](n, strs)

	// print
	fmt.Println(strs)
}

func sortArr[T int | string](n int, arr []T) (s []T) {
	// iterates the array
	for i := 0; i <= n; i++ {
		// for optimized if the array already sorted
		isSwap := false

		// the GIF explanation for this algorithm can
		// be found on this link https://commons.wikimedia.org/wiki/File:Bubble-sort-example-300px.gif
		// Basically, bubble-sort will compare the current and next index value
		// and swap if the current more than next index value
		for j := 0; j < n-i; j++ {
			// > operator means ascending sort
			// < operator means descending sort
			if arr[j] > arr[j+1] {
				isSwap = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		// break if not swapped
		if !isSwap {
			break
		}
	}

	return
}
