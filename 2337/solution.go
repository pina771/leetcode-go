package main

import (
	"fmt"
)

func main() {
	fmt.Println(canChange("_L__R__RL", "L_____RLR"))
}

// NOTE: Odustao od rješenja, ne ide baš
func canChange(start string, target string) bool {
	i, j := 0, 0
	n := len(start)
	for i < n && j < n {
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}

		if i == n && j == n {
			return true
		}

		if i == n || j == n || start[i] != target[j] {
			return false
		}

		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++

		j++
	}
	return true
}
