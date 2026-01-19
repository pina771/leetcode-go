package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	fmt.Println(tupleSameProduct(a))
}

func tupleSameProduct(nums []int) int {
	retval := 0
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			m[nums[i]*nums[j]]++
		}
	}
	fmt.Printf("%v", m)

	for _, v := range m {
		numOfPairs := v / 2

		// for the number of combinations - see visualization in notebook
		if numOfPairs > 1 {
			numOfCombinations := (numOfPairs - 1) * (numOfPairs) / 2
			tuplesForCombinations := 8 * numOfCombinations
			retval += tuplesForCombinations
		}
	}

	return retval
	// naive implementation
	// retval := 0
	// for a := 0; a < len(nums); a++ {
	// 	for b := 0; b < len(nums); b++ {
	// 		if b == a {
	// 			continue
	// 		}
	// 		ab := nums[a] * nums[b]
	//
	// 		for c := 0; c < len(nums); c++ {
	// 			if c == a || c == b {
	// 				continue
	// 			}
	// 			for d := 0; d < len(nums); d++ {
	// 				if d == a || d == b || d == c {
	// 					continue
	// 				}
	// 				cd := nums[c] * nums[d]
	// 				if ab == cd {
	// 					fmt.Printf("[%d, %d, %d, %d]\n",
	// 						nums[a], nums[b], nums[c],
	// 						nums[d],
	// 					)
	// 					retval++
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// return retval
}
