package main

// 233 https://leetcode-cn.com/problems/number-of-digit-one/
// 模拟
func countDigitOne(n int) int {
	idx, ans, pre := 1, 0, 0
	for n > 0 {
		tmp := n % 10
		n /= 10
		if tmp > 1 {
			ans += (n + 1) * idx
		} else if tmp == 1 {
			ans += n*idx + pre + 1
		} else {
			ans += n * idx
		}
		pre += tmp * idx
		idx *= 10
	}
	return ans
}
