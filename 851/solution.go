package main

import (
	"fmt"
)

func main() {
	richer := [][]int{
		{1, 0},
		{2, 1},
		{3, 1},
		{3, 7},
		{4, 3},
		{5, 3},
		{6, 3},
	}
	quiet := []int{3, 2, 5, 4, 6, 1, 7, 0}

	fmt.Println(loudAndRich(richer, quiet))
}

func loudAndRich(richer [][]int, quiet []int) []int {
	richerThan := make(map[int][]int)

	for idx := range quiet {
		richerThan[idx] = make([]int, 0)
	}

	for _, val := range richer {
		lessThan := val[1]
		richerThan[lessThan] = append(richerThan[lessThan], val[0])
	}

	answer := make([]int, len(quiet))
	for i := 0; i < len(answer); i++ {
		answer[i] = -1
	}
	for i := 0; i < len(answer); i++ {
		answer[i] = dfs(richerThan, quiet, i, answer)
	}

	return answer
}

func dfs(richerThan map[int][]int, quiet []int, idx int, answer []int) int {
	// not cached
	if answer[idx] == -1 {
		answer[idx] = idx
		for _, val := range richerThan[idx] {
			candidate := dfs(richerThan, quiet, val, answer)
			if quiet[candidate] < quiet[answer[idx]] {
				answer[idx] = candidate
			}
		}
	}

	// cached
	return answer[idx]
}
