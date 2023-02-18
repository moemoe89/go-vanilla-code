package main

import "fmt"

func main() {
	// sort int
	// prepare the int array
	nums := []int{6, 9, 1, 2, 13, 66, 3, 4, 9}

	left := 0
	right := len(nums) - 1

	// sort
	quickSort[int](nums, left, right)

	// print
	fmt.Println(nums)

	// sort string
	// prepare the int array
	strs := []string{"vcfgdf", "asdasda", "123asfdf", "xtyrhrt", "dtfgdfgh"}

	left = 0
	right = len(strs) - 1

	// sort
	quickSort[string](strs, left, right)

	// print
	fmt.Println(strs)
}

// the GIF explanation for this algorithm can
// be found on this link https://commons.wikimedia.org/wiki/File:Quicksort-example.gif
func quickSort[T int | string](arr []T, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition[T int | string](arr []T, left int, right int) int {
	pivot := arr[right]

	i := left

	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}
