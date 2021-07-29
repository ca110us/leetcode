package main

import "math"

// 171 https://leetcode-cn.com/problems/excel-sheet-column-number/
// 和 2 机制转 10 进制同理 这里是 26 进制
// AB => 1*26^1 + 2*26^0
func titleToNumber(columnTitle string) (ans int) {
	maxIndex := len(columnTitle)
	for i := 0; i < maxIndex; i++ {
		ans += (int(columnTitle[maxIndex-1-i]) - 64) * int(math.Pow(26, float64(i)))
	}

	return
}
