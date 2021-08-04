package main

import "sort"

// 611 https://leetcode-cn.com/problems/valid-triangle-number/
// 先排序 然后先找最大的边
// 接下来用双指针的方法去找次大边和最小边
// 每次确定一个次大边 就找到最小边最小的 idx 然后计算多少种答案
func triangleNumber(nums []int) (ans int) {
	ln := len(nums)

	sort.Ints(nums)
	for i := ln - 1; i > 0; i-- {
		// k 越小 j 只能越大
		for k, j := i-1, 0; j < k; k-- {
			for {
				if j < k && nums[k]+nums[j] <= nums[i] {
					j++
				} else {
					break
				}
			}
			ans += k - j
		}

	}

	return ans
}
