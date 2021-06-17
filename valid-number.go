package main

import "strings"

// 65 https://leetcode-cn.com/problems/valid-number/
// 状态机 图 —— https://pic.leetcode-cn.com/0ae239f74ce6ecaaf7c9044291b3fcdc8c5e60ac28dc447b7712a1500f9b3e01-1.png
func isNumber(s string) bool {
	dfa := []map[string]int{
		{},
		{"blank": 1, "sign": 2, "digit": 3, ".": 4},
		{"digit": 3, ".": 4},
		{"digit": 3, ".": 5, "e": 6, "blank": 9},
		{"digit": 5},
		{"digit": 5, "e": 6, "blank": 9},
		{"sign": 7, "digit": 8},
		{"digit": 8},
		{"digit": 8, "blank": 9},
		{"blank": 9},
	}

	curr_state := 1
	s = strings.TrimRight(s, " ")
	for i := 0; i < len(s); i++ {
		c := ""
		switch s[i] {
		case ' ':
			c = "blank"
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			c = "digit"
		case '+', '-':
			c = "sign"
		case '.':
			c = "."
		case 'e', 'E':
			c = "e"
		default:
			return false
		}
		if _, ok := dfa[curr_state][c]; !ok {
			return false
		}
		curr_state = dfa[curr_state][c]
	}

	return curr_state == 3 || curr_state == 5 || curr_state == 8 || curr_state == 9
}
