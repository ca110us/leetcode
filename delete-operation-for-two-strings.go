package main

// 583 https://leetcode-cn.com/problems/delete-operation-for-two-strings/
// 动态规划
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for k, _ := range dp {
		dp[k] = make([]int, len(word2)+1)
		dp[k][0] = k
	}
	for k, _ := range dp[0] {
		dp[0][k] = k
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+2, min(dp[i-1][j], dp[i][j-1])+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
