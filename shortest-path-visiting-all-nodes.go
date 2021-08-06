package main

// 210 https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/4
// bfs + 状态压缩 起点开始就在 0 处是一种状态
// 从 0 走到 1 再回到 0，虽然都是在 0 但是是两种状态
// 如果该状态出现过了 就不再尝试（见下）

// 这题和一般广度优先搜索不一样的点是，会走遍所有的点，并且可以会重复到达同一个点
// 所以当回到同一个点的时候，要判断这时的状态和之前在这个点时候的状态是否一样，如果一样就没有必要再次尝试了
func shortestPathLength(graph [][]int) int {
	ln := len(graph)
	type node struct{ idx, mask, dist int }
	q := []node{}
	vis := make([][]bool, ln)
	for i := 0; i < ln; i++ {
		q = append(q, node{i, 1 << i, 0})
		vis[i] = make([]bool, 1<<ln)
		vis[i][1<<i] = true
	}

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if n.mask == (1<<ln)-1 {
			return n.dist
		}

		for _, v := range graph[n.idx] {
			newMask := n.mask | 1<<v
			if !vis[v][newMask] {
				q = append(q, node{v, newMask, n.dist + 1})
				vis[v][newMask] = true
			}
		}
	}
	return 0
}
