package main

// 552 https://leetcode-cn.com/problems/student-attendance-record-ii/
// dp 动态规划
func checkRecord(n int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		ndp := make([][]int, 2)
		for j := 0; j < 2; j++ {
			ndp[j] = make([]int, 3)
		}
		ndp[0][0] = (dp[0][1] + dp[0][2] + dp[0][0]) % (1e9 + 7)
		ndp[0][1] = dp[0][0]
		ndp[0][2] = dp[0][1]
		ndp[1][0] = (dp[0][1] + dp[0][2] + dp[0][0] + dp[1][1] + dp[1][2] + dp[1][0]) % (1e9 + 7)
		ndp[1][1] = dp[1][0]
		ndp[1][2] = dp[1][1]
		dp = ndp
	}

	return (dp[0][0] + dp[0][1] + dp[0][2] + dp[1][0] + dp[1][1] + dp[1][2]) % (1e9 + 7)
}
