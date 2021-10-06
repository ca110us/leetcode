package main

import "math"

// 414 https://leetcode-cn.com/problems/third-maximum-number/
// 三个变量
func thirdMax(nums []int) int {
	var signedNum int
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num == max1 || num == max2 || num == max3 {
			continue
		}

		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
			signedNum++
		} else if num > max2 {
			max3 = max2
			max2 = num
			signedNum++
		} else if num > max3 {
			max3 = num
			signedNum++
		}
	}
	if signedNum < 3 { //赋值次数小于3次，则表示只有3个以内的数字，即需要返回最大值
		return max1
	}

	return max3
}
