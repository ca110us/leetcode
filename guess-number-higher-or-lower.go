package main

// 374 https://leetcode-cn.com/problems/guess-number-higher-or-lower/submissions/
// 二分查找
func guess(num int) int

func guessNumber(n int) int {
	start := 1
	end := n

	for {
		mid := (start + end) / 2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		case 0:
			return mid
		}
	}
}
