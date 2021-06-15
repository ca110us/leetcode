package main

// 852 https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/
// 二分查找
func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
