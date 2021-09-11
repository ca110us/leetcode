package main

// 600 https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/
// 动态规划 字典树
func findIntegers(n int) int {
	x := 1
	ans := uint32((n & 1) + 1)
	for i := 1; i < 32; i++ {
		x <<= 1
		if x > n {
			break
		}
		if x&n != 0 {
			if (x>>1)&n != 0 {
				ans = g[i]
			} else {
				ans += g[i-1]
			}
		}
	}
	return int(ans)
}

var g [32]uint32

func init() {
	N := 32
	g[0] = 2
	g[1] = 3
	for i := 2; i < N; i++ {
		g[i] = g[i-1] + g[i-2]
	}
}
