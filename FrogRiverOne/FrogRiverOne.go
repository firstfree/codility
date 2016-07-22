package main

import "fmt"

func Solution(X int, A []int) int {
	steps, visited := X, make([]bool, X+1)

	for i, v := range A {
		if !visited[v] {
			visited[v] = true
			steps--
		}

		if steps == 0 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4}))
}
