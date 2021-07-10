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
