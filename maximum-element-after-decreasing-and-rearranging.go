package main

import "sort"

// 1846 https://leetcode-cn.com/problems/maximum-element-after-decreasing-and-rearranging/
// 先排序 再贪心
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)

	arr[0] = 1
	ln := len(arr)
	for i := 1; i < ln; i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}

	return arr[ln-1]
}
