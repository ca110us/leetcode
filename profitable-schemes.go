package main

// 879 https://leetcode-cn.com/problems/profitable-schemes/
// 背包问题，需要注意的是，这里定义的状态是工作利润至少为 k 的盈利计划总数
// dp[i][0] = 1 表示什么任务都没有的时候，利润至少为 0 的情况只有一种，就是什么任务都没有，无论人数是多少
// 倒序遍历时候有可能为负数，因为利润一定是非负数，所以取 0 时候的方案数一样
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}

	mod := 1e9 + 7
	for i, num := range group {
		p := profit[i]
		for j := n; j > num-1; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] = (dp[j][k] + dp[j-num][max(0, k-p)]) % int(mod)
			}
		}
	}

	return dp[n][minProfit]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
