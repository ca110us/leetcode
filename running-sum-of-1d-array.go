package main

// 1480 https://leetcode-cn.com/problems/running-sum-of-1d-array/
// 模拟
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}
