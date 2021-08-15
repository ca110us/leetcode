package main

// 576 https://leetcode-cn.com/problems/out-of-boundary-paths/
// dfs
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	// dp[i][j][k] 表示球移动 k 次之后位于坐标 (i,j) 的路径数量
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j int, mov int) int
	dfs = func(i, j int, mov int) int {
		if i < 0 || j < 0 || i > m-1 || j > n-1 {
			return 1
		}

		if mov == 0 {
			return 0
		}

		if dp[i][j][mov] != -1 {
			return dp[i][j][mov]
		}

		dp[i][j][mov] = (dfs(i-1, j, mov-1) + dfs(i+1, j, mov-1) + dfs(i, j-1, mov-1) + dfs(i, j+1, mov-1)) % (1e9 + 7)
		return dp[i][j][mov]
	}

	return dfs(startRow, startColumn, maxMove)
}
