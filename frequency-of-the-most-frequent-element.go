package main

import "sort"

// 1838 https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element/
// 滑动窗口 图：https://pic.leetcode-cn.com/1626652746-HaJjEU-SLB)HU1WR6%5BOM%5DSSKU382%7BT.png
// 阴影部分即为 k
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	ln := len(nums)
	max := 1
	sum := 0
	l := 0

	for r := 1; r < ln; r++ {
		sum += (nums[r] - nums[r-1]) * (r - l)
		for sum > k {
			sum -= nums[r] - nums[l]
			l++
		}

		if max < r-l+1 {
			max = r - l + 1
		}
	}

	return max
}
