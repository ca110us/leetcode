package main

// 66 https://leetcode-cn.com/problems/plus-one/
// 从后往前遍历
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	//全部为9
	ans := make([]int, len(digits)+1)
	ans[0] = 1
	for i := 1; i < len(ans)-1; i++ {
		ans[i] = 0
	}

	return ans
}
