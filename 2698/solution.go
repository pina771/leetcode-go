package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	retval := checkNumber(0, "100", 10)
	fmt.Println(retval)
	fmt.Println(punishmentNumber(37))
}

func punishmentNumber(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		iSq := i * i
		numberStr := strconv.Itoa(iSq)
		if checkNumber(0, numberStr, i) {
			sum += iSq
		}
	}
	return sum
}

func checkNumber(sum int, str string, num int) bool {
	if num == 1 {
		return true
	}
	for i := 1; i < len(str); i++ {
		left := str[:i]
		right := str[i:]

		leftInt, err := strconv.Atoi(left)
		if err != nil {
			fmt.Println("error during atoi", err)
			os.Exit(1)
		}
		rightInt, err := strconv.Atoi(right)
		if err != nil {
			fmt.Println("error during atoi", err)
			os.Exit(1)
		}

		if sum+leftInt+rightInt == num {
			return true
		}

		deeperVal := checkNumber(sum+leftInt, right, num)
		if deeperVal == true {
			return true
		}
	}

	return false
}
