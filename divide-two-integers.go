package main

// 29 https://leetcode-cn.com/problems/divide-two-integers/
// 位运算
func divide(dividend int, divisor int) int {
	bool := true // 记录结果是否是正数
	if dividend < 0 {
		bool = !bool
		dividend = -dividend
	}
	if divisor < 0 {
		bool = !bool
		divisor = -divisor
	}
	res := 0
	for dividend > divisor { // 循环减
		count := 0
		temp := divisor // 临时变量
		for temp < dividend {
			temp = temp << 1 // 位运算
			count++
		}
		temp = temp >> 1
		res += 1 << (count - 1)
		dividend = dividend - temp // 从被除数中减去临时变量
	}
	if dividend == divisor {
		res++
	}
	if !bool {
		res = -res
	}
	if (1<<31 <= res) || (res < (-1 << 31)) {
		return 1<<31 - 1
	} // 处理超界
	return res
}
