package main

import "fmt"

func main() {
	fmt.Println(minSteps(12))
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	fmt.Println("minSteps n=", n)

	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			numRepeated := n / i
			return numRepeated + minSteps(i)
		}
	}

	return n
}
