package main

import (
	"container/heap"
	"sort"
)

// 1337 https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
// 和昨天的 987 数据结构一样 主要是自定义排序
func kWeakestRows(mat [][]int, k int) []int {
	h := hp{}
	for i, row := range mat {
		v := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		//h = append(h, pair{v, i})
		heap.Push(&h, pair{v, i})
	}
	//heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).index
	}
	return ans
}

// 也可以这样
// func kWeakestRows(mat [][]int, k int) []int {
// 	h := hp{}
// 	for i, row := range mat {
// 		v := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
// 		heap.Push(&h,pair{v,i})
// 	}

// 	ans := make([]int, k)
// 	for i := range ans {
// 		ans[i] = heap.Pop(&h).(pair).index
// 	}
// 	return ans
// }

type pair struct{ val, index int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.val < b.val || a.val == b.val && a.index < b.index
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
