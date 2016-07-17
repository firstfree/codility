package main

import (
	"fmt"
	"math"
)

func Solution(N int) int {
	startBit, len := -1, 0

	for N > 0 {
		k := N & -N

		endBit := int(math.Log2(float64(k)))

		if (startBit != -1) && ((endBit - startBit - 1) > len) {
			len = endBit - startBit - 1
		}

		startBit = endBit

		N = N & (N - 1)
	}

	return len
}

func main() {
	fmt.Println(Solution(1041))
}
