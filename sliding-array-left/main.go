package main

import "fmt"

func slidingArrayLeft(k int, arr []int) []int {
	return append(arr[k:], arr[:k]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	arr = slidingArrayLeft(3, arr)

	fmt.Println(arr)
}
