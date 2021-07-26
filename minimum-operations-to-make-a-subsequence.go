package main

import "sort"

// 1713 https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence
// 二分 + 动态规划
func minOperations(target []int, arr []int) int {
	m := make(map[int]int)
	for i, v := range target {
		m[v] = i
	}

	dp := make([]int, 0, len(arr))
	for _, v := range arr {
		// 求最大上升自序列
		if p, e := m[v]; e {
			if len(dp) == 0 || p > dp[len(dp)-1] {
				dp = append(dp, p)
			} else {
				if idx := sort.SearchInts(dp, p); idx < len(dp) {
					dp[idx] = p
				}
			}
		}
	}

	return len(target) - len(dp)
}
