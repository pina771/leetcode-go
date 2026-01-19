package main

import "fmt"

func main() {
	a := []int{3, 1}
	// b := []int{3, 5, 1}
	res := search(a, 1)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		fmt.Println("left=", left)
		fmt.Println("right=", right)
		fmt.Println("mid=", mid)
		fmt.Println("=======")

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			if nums[mid] > nums[right] && nums[right] >= target {
				fmt.Println("Here1")
				left = mid + 1
			} else {
				fmt.Println("Here2")
				right = mid - 1
			}
		} else if nums[mid] < target {
			if nums[mid] < nums[left] && nums[left] <= target {
				fmt.Println("Here3")
				right = mid - 1
			} else {
				fmt.Println("Here4")
				left = mid + 1
			}
		}
	}

	return -1
}
