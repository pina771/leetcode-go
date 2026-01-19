package main

import "testing"

func Test1(t *testing.T) {
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
	expOut := []int{5, 5, 2, 5, 4, 5, 6, 7}

	actualOut := loudAndRich(richer, quiet)

	for idx := range expOut {
		if expOut[idx] != actualOut[idx] {
			t.Fatalf("test failed, expected=%v , got=%v", expOut, actualOut)
		}
	}
}
