package main

// 1137 https://leetcode-cn.com/problems/n-th-tribonacci-number/
// 动态规划
func tribonacci(n int) int {
	if n <= 2 {
		if n == 0 {
			return 0
		}
		return 1
	}
	a, b, c, d := 0, 0, 1, 1
	for i := 2; i < n; i++ {
		a = b
		b = c
		c = d
		d = a + b + c
	}
	return d
}
