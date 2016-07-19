package main

import "fmt"

func Solution(A []int) int {
	res := 0

	for _, v := range A {
		res ^= v
	}

	return res
}

func main() {
	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 3, 3, 9}))
}
