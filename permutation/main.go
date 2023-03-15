package main

import "fmt"

func permutation(arr []int) (result [][]int) {
	backtrack(arr, &result, 0)

	return result
}

func backtrack(arr []int, result *[][]int, start int) {
	n := len(arr)

	if start == n {
		tmp := make([]int, n)
		copy(tmp, arr)
		*result = append(*result, tmp)

		return
	}

	for i := start; i < n; i++ {
		arr[start], arr[i] = arr[i], arr[start]

		backtrack(arr, result, start+1)

		arr[start], arr[i] = arr[i], arr[start]
	}
}

func main() {
	fmt.Println(permutation([]int{1, 2}))
}
