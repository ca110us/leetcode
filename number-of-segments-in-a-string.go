package main

// 434 https://leetcode-cn.com/problems/number-of-segments-in-a-string/
// 模拟
func countSegments(s string) (ans int) {
	for i, ch := range s {
		if (i == 0 || s[i-1] == ' ') && ch != ' ' {
			ans++
		}
	}
	return
}
