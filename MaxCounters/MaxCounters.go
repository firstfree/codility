package main

import "fmt"

func Solution(N int, A []int) []int {
	res := make([]int, N)
	currentMax, lastMax := 0, 0

	for _, v := range A {
		if v < N+1 {
			if res[v-1] < lastMax {
				res[v-1] = lastMax
			}

			res[v-1]++

			if currentMax < res[v-1] {
				currentMax = res[v-1]
			}
		} else {
			lastMax = currentMax
		}
	}

	for i, v := range res {
		if v < lastMax {
			res[i] = lastMax
		}
	}

	return res
}

func main() {
	fmt.Println(Solution(5, []int{3, 4, 4, 6, 1, 4, 4}))
	fmt.Println(Solution(2, []int{1, 1, 1, 3}))
	fmt.Println(Solution(10, []int{1, 11, 11, 11, 2, 11, 2, 1, 2, 2, 3, 4, 4, 3}))
}
