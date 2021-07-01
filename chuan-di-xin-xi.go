package main

// LCP 07 https://leetcode-cn.com/problems/chuan-di-xin-xi/
// DFS 每一次探索到 k 层就结束（要求 k 层刚好为 n-1）
func numWays(n int, relation [][]int, k int) int {
	m := make(map[int][]int)

	for _, v := range relation {
		from := v[0]
		to := v[1]

		if _, ok := m[from]; !ok {
			m[from] = []int{to}
		} else {
			m[from] = append(m[from], to)
		}
	}

	var dfs func(now, sum int)
	ans := 0
	dfs = func(now, sum int) {
		if sum == k {
			if now == n-1 {
				ans++
			}
			return
		}

		if _, ok := m[now]; !ok {
			return
		}

		for _, nxt := range m[now] {
			dfs(nxt, sum+1)
		}
	}

	dfs(0, 0)
	return ans
}
