package main

// https://leetcode-cn.com/problems/continuous-subarray-sum/
// 前缀和+哈希表
func checkSubarraySum(nums []int, k int) bool {
	m := map[int]int{0: -1}
	sum := 0
	for i, v := range nums {
		sum += v
		rem := sum % k
		if j, ok := m[rem]; ok {
			if i-j > 1 {
				return true
			}
		} else {
			m[rem] = i
		}
	}
	return false
}
