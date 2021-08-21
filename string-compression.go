package main

import (
	"strconv"
)

// 443 https://leetcode-cn.com/problems/string-compression/
// 双指针
func compress(chars []byte) int {
	l, r, idx := 0, 0, 0
	for r <= len(chars)-1 {
		now := chars[l]
		for chars[r] == now {
			r++
			if r >= len(chars) {
				break
			}
		}

		weight := r - l
		if weight == 1 {
			chars[idx] = now
			idx++
		} else {
			numStr := strconv.Itoa(weight)
			chars[idx] = now
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
