package main

import (
	"container/heap"
	"sort"
)

// 502 https://leetcode-cn.com/problems/ipo/
// 根堆 + 贪心
type node struct {
	profit, need int
}

type hp []node

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	if h[i].profit > h[j].profit {
		return true
	}
	return false
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Pop() (v interface{}) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(node))
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := make([]node, 0)
	for i := 0; i < len(profits); i++ {
		n = append(n, node{profit: profits[i], need: capital[i]})
	}
	sort.Slice(n, func(i, j int) bool {
		if n[i].need < n[j].need {
			return true
		}
		return false
	})
	h := make(hp, 0)
	for idx := 0; k > 0; k-- {
		for idx < len(n) && w >= n[idx].need {
			heap.Push(&h, n[idx])
			idx++
		}

		if h.Len() != 0 {
			w += heap.Pop(&h).(node).profit
		}

	}
	return w
}
