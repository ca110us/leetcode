package main

// 525 https://leetcode-cn.com/problems/contiguous-array/submissions/
// 计算特殊前缀和，是零则减1，是1则加1，hash记录前缀和数值第一次出现的 index
// 如果前缀和为零则当前i即为最长长度，否则当前i减第一次出现的 index 即为最长长度
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

// func main() {
// 	findMaxLength([]int{0, 1, 1, 0, 0, 0, 1})
// }
