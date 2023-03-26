package main

import "fmt"

func pow(a, b int64) int64 {
	result := int64(1)

	for i := int64(0); i < b; i++ {
		result *= a
	}

	return result
}

func main() {
	fmt.Println(pow(2, 0)) // 2^0
	fmt.Println(pow(2, 1)) // 2^1
	fmt.Println(pow(2, 2)) // 2^2
}
