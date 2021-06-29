package main

import "fmt"

// 168 https://leetcode-cn.com/problems/excel-sheet-column-title/
// 因为是从 1 开始，所以先减 1 让整体偏移，然后参照 10 进制转 16 进制的算法
func convertToTitle(columnNumber int) string {
	str := ""
	for columnNumber > 0 {
		columnNumber--
		str = fmt.Sprintf("%s%s", string(rune(columnNumber%26+65)), str)
		columnNumber /= 26
	}
	return str
}
