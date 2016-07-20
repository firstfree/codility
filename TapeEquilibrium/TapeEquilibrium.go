package main

import (
	"fmt"
	"math"
)

func Solution(A []int) int {
	sum := 0
	left := A[0]

	for i := 1; i < len(A); i++ {
		sum += A[i]
	}

	res := int(math.Abs(float64(sum - left)))

	for i := 1; i < len(A)-1; i++ {
		left += A[i]
		sum -= A[i]

		diff := int(math.Abs(float64(sum - left)))

		if diff < res {
			res = diff
		}
	}

	return res
}

func main() {
	fmt.Println(Solution([]int{-1, -2, -3, -4, -5, -6}))
}
