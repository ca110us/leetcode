package main

import "sort"

// 441 https://leetcode-cn.com/problems/arranging-coins/
// 二分查找
func arrangeCoins(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}
