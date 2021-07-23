package main

import (
	"strconv"
)

// 1736 https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits/
// 列举所有情况即可
func maximumTime(time string) string {
	ba := []byte(time)
	for i, v := range ba {
		if string(v) == "?" {
			switch i {
			case 0:
				if string(time[1]) == "?" {
					ba[i] = byte('2')
					continue
				}
				if nxt, _ := strconv.ParseInt(string(time[1]), 10, 64); nxt < 4 {
					ba[i] = byte('2')
				} else {
					ba[i] = byte('1')
				}
			case 1:
				if string(time[0]) == "?" {
					ba[i] = byte('3')
					continue
				}
				if pre, _ := strconv.ParseInt(string(time[0]), 10, 64); pre < 2 {
					ba[i] = byte('9')
				} else {
					ba[i] = byte('3')
				}
			case 3:
				ba[i] = byte('5')
			case 4:
				ba[i] = byte('9')
			}
		}
	}

	return string(ba)
}
