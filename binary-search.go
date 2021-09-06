package main

// 704 https://leetcode-cn.com/problems/binary-search/
// 二分查找
func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
