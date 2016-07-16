package main

import (
	"fmt"
	"math"
)

func Solution(N int) int {
	left, len := -1, 0

	for N > 0 {
		k := N & -N

		right := int(math.Log2(float64(k)))

		if (left != -1) && ((right - left - 1) > len) {
			len = right - left - 1
		}

		left = right

		N = N & (N - 1)
	}

	return len
}

func main() {
	fmt.Println(Solution(1041))
}
