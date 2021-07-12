package main

import "sort"

// 274 https://leetcode-cn.com/problems/h-index/
// 先 google 一下什么是 H 指数吧 ...
func hIndex(citations []int) int {
	sort.Ints(citations)

	ans := 0
	for i := len(citations) - 1; i >= 0 && citations[i] > ans; i-- {
		ans++
	}

	return ans
}

// 275 https://leetcode-cn.com/problems/h-index/
// 第二种方法 二分查找
func hIndex2(citations []int) int {
	left, right := 0, len(citations)-1
	for left < right {
		mid := (left + right) / 2

		if citations[mid] < len(citations)-mid {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if citations[left] == 0 {
		return 0
	}

	return len(citations) - left
}
