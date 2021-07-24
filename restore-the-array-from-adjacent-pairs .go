package main

// 1743 https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs/
// 哈希 选定 1 个只有一个相邻元素的端点，朝一个方向递推
func restoreArray(adjacentPairs [][]int) []int {
	ln := len(adjacentPairs) + 1
	m := make(map[int][]int, ln)

	for _, v := range adjacentPairs {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}

	ans := make([]int, ln)
	for k, v := range m {
		if len(v) == 1 {
			ans[0] = k
			ans[1] = v[0]
			break
		}
	}

	// 可以直接用加减法 而不用去判断
	for i := 2; i < ln; i++ {
		pre := ans[i-1]
		ans[i] = m[pre][0] + m[pre][1] - ans[i-2]
	}

	return ans
}
