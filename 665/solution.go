package main

import "fmt"

func main() {
	a := []int{1, 5, 2, 4, 3}
	fmt.Println(checkPossibility(a))
}

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			cnt++

			if i-2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return cnt <= 1
}
