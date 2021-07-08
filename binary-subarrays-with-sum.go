package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum/
// 前缀和 + hash 表
func numSubarraysWithSum(nums []int, goal int) int {
	m := make(map[int]int, 0)
	ans := 0

	sum := 0
	for _, n := range nums {
		// 先统计前缀和，为了加上默认和为 0 的 1 种情况
		m[sum]++

		sum += n
		if sum >= goal {
			ans += m[sum-goal]
		}
	}

	return ans
}
