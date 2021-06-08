package main

// 1049 https://leetcode-cn.com/problems/last-stone-weight-ii/
// 将石头分为两堆，相互碰撞;让其中一个堆的重量和尽量接近质量总和的一半，就是背包问题
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, w := range stones {
		sum += w
	}

	avg := sum >> 1

	dp := make([]int, avg+1)
	dp[0] = 0

	for _, w := range stones {
		for j := avg; j > w-1; j-- {
			dp[j] = max(dp[j], dp[j-w]+w)
		}
	}

	max := dp[avg]
	less := sum - max

	min := max - less
	if min < 0 {
		min = -min
	}

	return min
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
