package main

import "fmt"

func slidingArrayRight(k int, arr []int) []int {
	n := len(arr)

	k = k % n

	return append(arr[n-k:], arr[:n-k]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	arr = slidingArrayRight(3, arr)

	fmt.Println(arr)
}
