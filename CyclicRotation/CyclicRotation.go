package main

import "fmt"

func Solution(A []int, K int) []int {
	len := len(A)
	res := make([]int, len)

	for i, v := range A {
		res[(i+K)%len] = v
	}

	return res
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4, 5, 6}, 3))
}
