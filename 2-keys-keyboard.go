package main

import "math"

// 650 https://leetcode-cn.com/problems/2-keys-keyboard/
// 动态规划
func minSteps(n int) int {
	h := int(math.Sqrt(float64(n)))
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j <= h; j++ {
			if i%j == 0 {
				// dp[i] = Min(dp[i], dp[j] + dp[i/j])
				dp[i] = dp[j] + dp[i/j]
			}
		}
	}
	return dp[n]
}

func Min(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}
