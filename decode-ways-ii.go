package main

// 639 https://leetcode-cn.com/problems/decode-ways-ii/
// 动态规划
func numDecodings(s string) int {
	a, b, c := 0, 1, 0
	m := int(1e9 + 7)
	for i := 1; i <= len(s); i++ {
		c = 0
		if s[i-1] == '*' {
			c = (c + b*9) % m
			if i > 1 && s[i-2] == '*' {
				c = (c + a*15) % m
			} else if i > 1 && s[i-2] == '1' {
				c = (c + a*9) % m
			} else if i > 1 && s[i-2] == '2' {
				c = (c + a*6) % m
			}
		} else if s[i-1] != '0' {
			c += b
			c %= m
			if i > 1 && s[i-2] == '*' {
				if s[i-1] <= '6' {
					c = (c + a*2) % m
				} else {
					c = (c + a) % m
				}
			} else if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
				c = (c + a) % m
			}
		} else {
			if i > 1 && s[i-2] == '*' {
				c = (c + a*2) % m
			} else if i > 1 && s[i-2] > '0' && s[i-2] <= '2' {
				c = (c + a) % m
			}
		}
		a, b = b, c
	}
	return c
}
