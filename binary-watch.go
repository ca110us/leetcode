package main

import (
	"fmt"
	"math/bits"
)

// 401 https://leetcode-cn.com/problems/binary-watch/
// 枚举所有 时和分
func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return ans
}
