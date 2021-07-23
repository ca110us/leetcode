package main

// 1893 https://leetcode-cn.com/problems/check-if-all-the-integers-in-a-range-are-covered/
// 位运算

// 6 5 4 3 2 1
// 1 1 1 1 1 1 表示 1-6 都有 （bm）
// 0 1 1 1 1 0 表示 2-5 都有 （check）
// 最后｜（或） 一下 如果 bm 不变，则说明都有
func isCovered(ranges [][]int, left int, right int) bool {
	bm := 0
	for _, r := range ranges {
		// right = 6
		// 1<<6 - 1 => 111111
		// left = 3
		// 1<<2 - 1 => 11(000011)
		// 111111 ^ 11 => 111100
		b := (1<<r[1] - 1) ^ (1<<(r[0]-1) - 1)
		bm |= b
	}

	check := (1<<right - 1) ^ (1<<(left-1) - 1)

	if bm|check == bm {
		return true
	}
	return false
}
