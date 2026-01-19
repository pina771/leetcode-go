package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 2
	nums := []int{7, 8, 8}
	fmt.Println(smallestRangeII(nums, k))
}

func smallestRangeII(nums []int, k int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := nums[n-1] - nums[0]

	for i := 0; i < n-1; i++ {
		high := max(nums[n-1]-k, nums[i]+k)
		low := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, high-low)
	}
	return ans
}
