package main

import "strings"

// 474 https://leetcode-cn.com/problems/ones-and-zeroes/
// 0-1 package https://www.youtube.com/watch?v=Vz6uJ7iQaN0
// 两个维度的背包问题，相当与一个物品有质量和体积，背包有最大载重和最大容量
func findMaxForm(strs []string, m int, n int) int {
	tabl := make([][]int, m+1)
	for i := range tabl {
		tabl[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		zeros := strings.Count(str, "0")
		ones := len(str) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				tabl[j][k] = max(tabl[j][k], tabl[j-zeros][k-ones]+1)
			}
		}
	}

	return tabl[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
