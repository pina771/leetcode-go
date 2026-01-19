package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minLength("ABFCACDB"))
}

func minLength(s string) int {
	removedAB := strings.Split(s, "AB")
	withoutABStr := strings.Join(removedAB, "")

	removedCD := strings.Split(withoutABStr, "CD")
	withoutCDStr := strings.Join(removedCD, "")

	if len(withoutCDStr) == len(s) {
		return len(withoutCDStr)
	}
	return minLength(withoutCDStr)
}
