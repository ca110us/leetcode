package main

// 518 https://leetcode-cn.com/problems/coin-change-2/
// 搞清楚遍历顺序，因为可以无限选取，所以直接正序遍历 ！
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// stupid !!
// func change(amount int, coins []int) int {
// 	dp := make([]int, amount+1)
// 	dp[0] = 1

// 	for _, coin := range coins {
// 		for j := amount; j > coin-1; j-- {
// 			k := 1
// 			for {
// 				if k*coin > j {
// 					break
// 				}

// 				dp[j] += dp[j-(k*coin)]
// 				k++
// 			}
// 		}
// 	}

// 	return dp[amount]
// }
