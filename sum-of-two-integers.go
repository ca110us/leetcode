package main

// 371 https://leetcode-cn.com/problems/sum-of-two-integers/
// 位运算
func getSum(a int, b int) int {
	for b != 0 {
		tmp := a ^ b
		b = (a & b) << 1
		a = tmp
	}
	return a
}
