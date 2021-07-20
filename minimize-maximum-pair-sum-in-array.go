package main

import "sort"

// 1877 https://leetcode-cn.com/problems/minimize-maximum-pair-sum-in-array/
// 大小搭配 尽量平均分配 然后求最大 双指针
func minPairSum(nums []int) int {
	sort.Ints(nums)

	l := 0
	r := len(nums) - 1
	ans := nums[l] + nums[r]
	for i, j := l+1, r-1; i < j; i, j = i+1, j-1 {
		if nums[i]+nums[j] > ans {
			ans = nums[i] + nums[j]
		}
	}

	return ans
}
