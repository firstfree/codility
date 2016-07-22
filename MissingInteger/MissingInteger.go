package main

import "fmt"

func Solution(A []int) int {
	len := len(A)
	count := make([]int, len+1)

	for _, v := range A {
		if (v <= len) && (v > 0) {
			count[v]++
		}
	}

	for i := 1; i < len+1; i++ {
		if count[i] == 0 {
			return i
		}
	}

	return len + 1
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4, 5, 6, 12})) // 5
	fmt.Println(Solution([]int{1}))                    // 2
	fmt.Println(Solution([]int{1, 2, 3, 4, 5}))        // 6
	fmt.Println(Solution([]int{4, 5, 6, 2}))           // 1
}
