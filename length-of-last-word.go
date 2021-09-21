package main

// 58 https://leetcode-cn.com/problems/length-of-last-word/
// 反向遍历
func lengthOfLastWord(s string) (ans int) {
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		i--
		ans++
	}
	return
}
