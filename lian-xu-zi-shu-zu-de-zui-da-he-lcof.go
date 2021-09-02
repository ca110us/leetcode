package main

// offer - 42 https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// 动态规划
func maxSubArray(nums []int) int {
	ans := nums[0]
	// 设 dp[i] 表示以 nums[i] 结尾元素的最大和
	// 如果前面加起来是负数 那么最大和就是 nums[i] 本身
	// 所以这里可以巧借 nums 充当 dp 数组
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		ans = max(ans, nums[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
