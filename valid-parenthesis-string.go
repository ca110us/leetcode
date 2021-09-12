package main

// 678 https://leetcode-cn.com/problems/valid-parenthesis-string/
// 模拟
func checkValidString(s string) bool {
	lower, high := 0, 0
	for _, c := range []byte(s) {
		if c == '*' {
			if lower != 0 {
				lower--
			}
			high++
		} else if c == '(' {
			lower++
			high++
		} else {
			if lower > 0 {
				lower--
			}
			high--
			if high < 0 {
				return false
			}
		}
	}
	return lower == 0 && high >= 0
}
