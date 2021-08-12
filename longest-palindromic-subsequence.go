package main

// 516 https://leetcode-cn.com/problems/longest-palindromic-subsequence/
// dp[i][j] 表示下标 i 到 j 之间最长的子回文序列
func longestPalindromeSubseq(s string) int {
	ln := len(s)
	dp := make([][]int, ln)

	for i := 0; i < ln; i++ {
		dp[i] = make([]int, ln)
	}

	for i := ln - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < ln; j++ {
			if s[i] == s[j] {
				// 如果相等则最长比没增加这两位前又增加了两位
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 不等的话从 i+1 -> j 或 i -> j-1 中选一个较长者
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][ln-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
