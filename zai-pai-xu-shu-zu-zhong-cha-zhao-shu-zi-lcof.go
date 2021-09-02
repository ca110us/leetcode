package main

import "sort"

// offer - 53 https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// 不讲武德
func search(nums []int, target int) (sum int) {
	for _, v := range nums {
		if v == target {
			sum++
		}
	}

	return sum
}

// 二分法
func search(nums []int, target int) (ans int) {
	sort.Ints(nums)
	// 先定位到位置
	idx := sort.Search(len(nums)-1, func(x int) bool { return nums[x] >= target })
	for idx < len(nums) {
		if nums[idx] == target {
			ans++
			idx++
		} else {
			break
		}

	}

	return ans
}
