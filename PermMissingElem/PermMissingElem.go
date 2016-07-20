package main

import "fmt"

func Solution(A []int) int {
	visited := make([]int, len(A)+1)

	for _, v := range A {
		visited[v-1] = 1
	}

	for i, v := range visited {
		if v == 0 {
			return i + 1
		}
	}

	return 0
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 5}))
}
