package main

// 525 https://leetcode-cn.com/problems/contiguous-array/submissions/
// 计算特殊前缀和，是 0 则减 1，是 1 则加 1，hash 记录前缀和数值第一次出现的 index
// 如果前缀和为 0 则当前 i 即为最长长度，否则当前 i 减第一次出现的 index 即为最长长度
func findMaxLength(nums []int) int {
	m := map[int]int{}
	sum := 0
	max := 0
	for i, v := range nums {
		if v == 0 {
			sum -= 1
		} else {
			sum += 1
		}
		if sum == 0 {
			max = i + 1
		} else {
			if j, ok := m[sum]; ok {
				len := i - j
				if len > max {
					max = len
				}
			} else {
				m[sum] = i
			}
		}

	}
	return max
}
