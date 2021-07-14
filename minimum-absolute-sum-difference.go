package main

import (
	"sort"
)

// 1818 https://leetcode-cn.com/problems/minimum-absolute-sum-difference/
// 二分查找
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	cp := append(sort.IntSlice(nil), nums1...)
	cp.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		idx := cp.Search(v)
		// 重点 cp[idx-1]<=v cp[idx] >= v
		// 判断是否超过最大 index 防止溢出
		if idx < n {
			maxn = max(maxn, diff-(cp[idx]-v))
		}
		// 判断是否可以 -1 防止溢出
		if idx > 0 {
			maxn = max(maxn, diff-(v-cp[idx-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
