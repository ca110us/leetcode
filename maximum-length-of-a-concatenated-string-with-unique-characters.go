package main

import "math/bits"

// 1239 https://leetcode-cn.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
// backtrace 回溯
func maxLength(arr []string) int {
	masks := make([]int, 0, len(arr))
	for _, w := range arr {
		mask := 0
		for _, c := range w {
			mask |= 1 << int(c-'a')
		}
		if len(w) == bits.OnesCount(uint(mask)) {
			masks = append(masks, mask)
		}
	}

	res := 0
	var backtrace func(cur, mask int)
	backtrace = func(cur, mask int) {
		res = max(res, bits.OnesCount(uint(mask)))
		for i := cur; i < len(masks); i++ {
			if masks[i]&mask != 0 {
				continue
			}
			backtrace(i+1, mask|masks[i])
		}
	}
	backtrace(0, 0)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
