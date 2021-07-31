package main

import (
	"container/heap"
)

// 987 https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree/
// 维护一个优先队列 map 的 key 是树的广度，那么示例 1 中 key 的范围就是 -1 到 2，而且都会存在
// map 的 v 是下面 node 的结构体 depth 是深度 v 是值
// 深度遍历 遍历左节点的时候广度-- 深度++ 遍历右节点的时候 广度++ 深度--
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct{ depth, v int }
type hp []node

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	if h[i].depth == h[j].depth {
		return h[i].v < h[j].v
	}
	return h[i].depth < h[j].depth
}

func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(node)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func verticalTraversal(root *TreeNode) (ans [][]int) {
	m := make(map[int]*hp, 0)
	minBreadth := 0
	maxBreadth := 0

	var dfs func(n *TreeNode, breadth, depth int)
	dfs = func(n *TreeNode, breadth, depth int) {
		if breadth < minBreadth {
			minBreadth = breadth
		}

		if breadth > maxBreadth {
			maxBreadth = breadth
		}

		if h, ok := m[breadth]; ok {
			heap.Push(h, node{depth: depth, v: n.Val})
		} else {
			h := new(hp)
			heap.Push(h, node{depth: depth, v: n.Val})

			m[breadth] = h
		}

		if n.Left != nil {
			b := breadth - 1
			d := depth + 1
			dfs(n.Left, b, d)
		}

		if n.Right != nil {
			b := breadth + 1
			d := depth + 1
			dfs(n.Right, b, d)
		}
	}

	dfs(root, 0, 0)
	for i := minBreadth; i <= maxBreadth; i++ {
		h := m[i]
		a := make([]int, 0)
		for {
			if h.Len() == 0 {
				break
			}
			v := heap.Pop(h)
			a = append(a, v.(node).v)
		}
		ans = append(ans, a)
	}

	return ans
}
