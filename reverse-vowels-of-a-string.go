package main

import "strings"

// 345 https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
// 双指针
func reverseVowels(s string) string {
	ln := len(s)

	str := []byte(s)
	check := "aeiouAEIOU"
	for i, j := 0, ln-1; i < j; {
		for i < ln && !strings.Contains(check, string(str[i])) {
			i++
		}
		for j > 0 && !strings.Contains(check, string(str[j])) {
			j--
		}
		if i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	return string(str)
}
