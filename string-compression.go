package main

import (
	"strconv"
)

// 443 https://leetcode-cn.com/problems/string-compression/
// 双指针
func compress(chars []byte) int {
	l, r, idx, ln := 0, 0, 0, len(chars)
	for r <= ln-1 {
		for chars[r] == chars[l] {
			r++
			if r >= ln {
				break
			}
		}

		if r-l == 1 {
			chars[idx] = chars[l]
			idx++
		} else {
			numStr := strconv.Itoa(r - l)
			chars[idx] = chars[l]
			idx++
			for _, v := range numStr {
				chars[idx] = byte(v)
				idx++
			}
		}
		l = r
	}
	return idx
}
