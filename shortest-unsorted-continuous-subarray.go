package main

import "sort"

// 581 https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
// 排序后 双指针 最近真的有点忙 .....
func findUnsortedSubarray(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return 0
	}

	nums2 := make([]int, ln)
	copy(nums2, nums)
	sort.Ints(nums2)

	var i, j int
	for i = 0; i < ln-1; i++ {
		if nums2[i] != nums[i] {
			break
		}
	}

	for j = ln - 1; j > 0; j-- {
		if nums2[j] != nums[j] {
			break
		}
	}

	res := j - i + 1
	if res > 0 {
		return res
	}

	return 0
}
