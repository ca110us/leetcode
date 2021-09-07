package main

// 1221 https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
// 贪心
func balancedStringSplit(s string) (ans int) {
	mark := 0
	for _, r := range s {
		if r == 'L' {
			mark++
		}
		if r == 'R' {
			mark--
		}
		if mark == 0 {
			ans++
		}
	}
	if ans == 0 {
		return 1
	}
	return
}
