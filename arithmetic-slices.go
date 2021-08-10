package main

// 413 https://leetcode-cn.com/problems/arithmetic-slices/
// 差分 + 计数
func numberOfArithmeticSlices(nums []int) (ans int) {
	if len(nums) == 1 {
		return
	}
	d, cnt := nums[0]-nums[1], 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == d {
			cnt++
		} else {
			d, cnt = nums[i-1]-nums[i], 0
		}
		ans += cnt
	}
	return
}
