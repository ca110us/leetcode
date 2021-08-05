package main

// 802 https://leetcode-cn.com/problems/find-eventual-safe-states/
// 新类型题 理解官方题解 三色法 https://leetcode-cn.com/problems/find-eventual-safe-states/solution/zhao-dao-zui-zhong-de-an-quan-zhuang-tai-yzfz/
func eventualSafeNodes(graph [][]int) (ans []int) {
	ln := len(graph)
	color := make([]int, ln)
	var safe func(int) bool
	safe = func(i int) bool {
		if color[i] > 0 {
			return color[i] == 2
		}
		color[i] = 1
		for _, v := range graph[i] {
			if !safe(v) {
				return false
			}
		}
		color[i] = 2
		return true
	}
	for i := 0; i < ln; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}
