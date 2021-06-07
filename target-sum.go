package main

// 494 https://leetcode-cn.com/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if (sum-target)%2 == 1 {
		return 0

	}

	if sum < target {
		return 0
	}

	neg := (sum - target) / 2

	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for x := neg; x > num-1; x-- {
			dp[x] += dp[x-num]
		}
	}

	return dp[neg]
}
