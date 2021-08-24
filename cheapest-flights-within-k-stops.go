package main

import (
	"fmt"
	"math"
)

// 787 https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
// 动态规划
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if n < 1 {
		return -1
	}
	matrix := make([][]int, n)
	save := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		save[i] = make([]int, K+1)
	}

	for i := 0; i < len(flights); i++ {
		s, e, v := flights[i][0], flights[i][1], flights[i][2]
		matrix[s][e] = v
	}

	res := math.MaxInt32
	queue := make([]int, 0, 2*n)
	queue = append(queue, src)
	step := 0

	for len(queue) > 0 && step <= K {

		preLen := len(queue)
		if preLen > n {
			fmt.Println(len(queue))
		}
		// fmt.Println("start ", s, "  ", queue)
		for j := 0; j < preLen; j++ {
			s := queue[0]
			queue = queue[1:]
			for i := 0; i < n; i++ {
				if matrix[s][i] <= 0 {
					continue
				}
				if step == 0 {
					save[i][step] = 0 + matrix[s][i]
				} else {
					s := save[s][step-1] + matrix[s][i]
					if s < save[i][step] || save[i][step] == 0 {
						save[i][step] = s
					} else {
						continue
					}
				}
				if save[i][step] > res {
					continue
				}

				if i == dst {
					res = save[i][step]
				}
				queue = append(queue, i)
			}
		}
		step++
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}
