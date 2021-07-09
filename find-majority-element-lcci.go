package main

// 面试题 17.10 https://leetcode-cn.com/problems/find-majority-element-lcci/submissions/
// 这题我只想到了用哈希的方法，但是空间并不是 O(1)
// 这里摩尔投票的关键点是：每次将两个不同的元素进行「抵消」，如果最后有元素剩余，则「可能」为元素个数大于总数一半的那个
func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count = 1
			continue
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	sum := 0
	for _, num := range nums {
		if num == candidate {
			sum++
		}
	}

	if sum*2 > len(nums) {
		return candidate
	}

	return -1
}
