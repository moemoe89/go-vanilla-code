package main

import "fmt"

func binarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			right = mid - 1
		} else if arr[mid] < x {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 5))
}
