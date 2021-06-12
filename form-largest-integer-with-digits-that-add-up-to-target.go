package main

// 1449 https://leetcode-cn.com/problems/form-largest-integer-with-digits-that-add-up-to-target/
// 完全背包问题，加上状态转移的回溯，可以用 0 和 -1 来区分数据是否是合法的
func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1)

	// -1 和 0 进行区分，避免数据错误
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for _, c := range cost {
		for j := c; j <= target; j++ {
			// 判断数据合法
			if dp[j-c] >= 0 {
				dp[j] = max(dp[j], dp[j-c]+1)
			}
		}
	}

	if dp[target] < 0 {
		return "0"
	}

	// 优先考虑从大的值转移来
	t := target
	ans := make([]byte, 0, dp[target])
	for i := 8; i >= 0; i-- {
		for c := cost[i]; c <= t && dp[t] == dp[t-c]+1; t -= c {
			ans = append(ans, byte(i+'1'))
		}
	}

	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
