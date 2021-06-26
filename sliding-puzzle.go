package main

import (
	"strconv"
	"strings"
)

// 773 https://leetcode-cn.com/problems/sliding-puzzle/
// BFS ， 刚开始用 752 的双向搜索去做，也可以，但是数据较小，没必要用双向搜索
func slidingPuzzle(board [][]int) int {
	// 从头开始搜索
	start := ""
	end := "123450"
	arr := append(board[0], board[1]...)
	for _, n := range arr {
		start += strconv.Itoa(n)
	}

	if start == end {
		return 0
	}

	q := []string{start}

	// 记录到这个 node 的步数
	visited := make(map[string]int)
	// 头节点步数为 0
	visited[start] = 0

	judge := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	t := -1
	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		step := visited[node]
		if node == end {
			return step
		}

		zeroIndex := strings.Index(node, "0")
		for i := 0; i < len(judge[zeroIndex]); i++ {
			for _, i := range judge[zeroIndex] {
				p := []rune(node)
				p[zeroIndex], p[i] = p[i], p[zeroIndex]

				pnode := string(p)
				// 已经遍历到过了
				if _, ok := visited[pnode]; ok {
					continue
				}

				visited[pnode] = step + 1
				q = append(q, pnode)
			}
		}
	}
	return t
}
