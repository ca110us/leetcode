package main

// 476 https://leetcode-cn.com/problems/number-complement/
// 位运算
func findComplement(num int) int {
	i := 1
	for i <= num {
		i <<= 1
	}
	return i - num - 1
}
