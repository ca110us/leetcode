package main

// 1894 https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk/submissions/
// 模拟
func chalkReplacer(chalk []int, k int) int {
	total := 0
	for _, num := range chalk {
		total += num
	}
	k %= total
	for i, need := range chalk {
		if k < need {
			return i
		}
		k -= need
	}
	return 0
}
