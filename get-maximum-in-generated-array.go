package main

// 1646 https://leetcode-cn.com/problems/get-maximum-in-generated-array/
// 根据题目 模拟
func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}

	res := make([]int, 512)
	res[0] = 0
	res[1] = 1
	max := 0
	for i := 1; i < n+1; i++ {
		res[2*i] = res[i]
		res[2*i+1] = res[i] + res[i+1]
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}
