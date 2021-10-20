package main

import "math"

// 453 https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
// 遍历
func minMoves(nums []int) int {
	s, m := 0, math.MaxInt32
	for _, num := range nums {
		s, m = s+num, min(m, num)
	}
	return s - m*len(nums)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
