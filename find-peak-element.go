package main

// 162 https://leetcode-cn.com/problems/find-peak-element/
// 二分查找
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if gt(nums, mid, mid-1) && gt(nums, mid, mid+1) {
			return mid
		} else if gt(nums, mid-1, mid) {
			right = mid - 1
		} else if gt(nums, mid+1, mid) {
			left = mid + 1
		}
	}
	return -1
}

func gt(nums []int, i, j int) bool {
	if i == -1 || i == len(nums) {
		return false
	}
	if j == len(nums) || j == -1 {
		return true
	}
	return nums[i] > nums[j]
}
