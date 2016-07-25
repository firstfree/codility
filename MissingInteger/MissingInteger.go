package main

import "fmt"

func Solution(A []int) int {
	len := len(A)
	visited := make([]bool, len)

	for _, v := range A {
		if (v > 0) && (v <= len) {
			visited[v-1] = true
		}
	}

	for i, v := range visited {
		if !v {
			return i + 1
		}
	}

	return len + 1
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4, 5, 6, 12})) // 7
	fmt.Println(Solution([]int{1}))                    // 2
	fmt.Println(Solution([]int{1, 2, 3, 4, 5}))        // 6
	fmt.Println(Solution([]int{4, 5, 6, 2}))           // 1
}
