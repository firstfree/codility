package main

import (
	"fmt"
	"math"
)

func Solution(X int, Y int, D int) int {
	return int(math.Ceil(float64(Y-X) / float64(D)))
}

func main() {
	fmt.Println(Solution(10, 85, 30))
}
