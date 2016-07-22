package main

import "fmt"

func Solution(A []int) int {
	visited := make([]bool, len(A))

	for _, v := range A {
		if (v > len(A)) || visited[v-1] {
			return 0
		}

		visited[v-1] = true
	}

	return 1
}

func main() {
	fmt.Println(Solution([]int{4, 1, 3, 2}))
}
