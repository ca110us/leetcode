package main

// Offer 10 https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 动态规划
func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
