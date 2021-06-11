package main

import "math"

// 279 https://leetcode-cn.com/problems/perfect-squares/

// 动态规划 dp[i] = x 表示 i 的完全平方数最小数量为 x
// func numSquares(n int) int {
// 	dp := make([]int, n+1)
// 	dp[0] = 0

// 	for i := 1; i <= n; i++ {
// 		dp[i] = i

// 		for j := 1; i-j*j >= 0; j++ {
// 			dp[i] = min(dp[i], dp[i-j*j]+1)
// 		}
// 	}
// 	return dp[n]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func is4km(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

// 数学方法 - 四平方和定理
// https://baike.baidu.com/item/%E5%9B%9B%E5%B9%B3%E6%96%B9%E5%92%8C%E5%AE%9A%E7%90%86
func numSquares(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if is4km(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
