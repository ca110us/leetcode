package main

// 526 https://leetcode-cn.com/problems/beautiful-arrangement/
// dfs 回溯
func countArrangement(n int) int {
	var ans int
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	var dfs func([]int, int)
	dfs = func(nums []int, i int) {
		if i == len(nums) {
			ans++
			return
		}

		for j := i; j < n; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			if nums[i]%(i+1) == 0 || (i+1)%nums[i] == 0 {
				dfs(nums, i+1)
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	dfs(nums, 0)
	return ans
}
