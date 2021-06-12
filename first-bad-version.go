package main

// 278 - https://leetcode-cn.com/problems/first-bad-version/submissions/
// 简单的二分查找
func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	start := 0
	end := n

	for start <= end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}
