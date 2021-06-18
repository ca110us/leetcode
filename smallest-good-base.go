package main

import (
	"math"
	"math/bits"
	"strconv"
)

// 483 https://leetcode-cn.com/problems/smallest-good-base/submissions/
// 数学 + 二分
func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	max := bits.Len(uint(num))
	for i := max - 1; i > 1; i-- {
		k := int(math.Pow(float64(num), 1/float64(i)))
		sum, s := 1, 1
		for j := 0; j < i; j++ {
			s *= k
			sum += s
		}
		if sum == num {
			return strconv.Itoa(k)
		}
	}

	return strconv.Itoa(num - 1)
}
