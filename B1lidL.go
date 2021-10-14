package main

import "sort"

// offer069 https://leetcode-cn.com/problems/B1IidL/
// 二分查找
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}
