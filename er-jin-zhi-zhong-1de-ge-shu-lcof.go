package main

// 剑指 offer - 15 https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
// 二进制中 1 的个数
// 位运算 n & (n - 1) 每次消除最右边的 1
func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		num &= num - 1
		sum++
	}

	return sum
}
