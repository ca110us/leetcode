package main

// offer53 https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// 不讲武德
func search(nums []int, target int) (sum int) {
	for _, v := range nums {
		if v == target {
			sum++
		}
	}

	return sum
}
