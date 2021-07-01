package main

import "sort"

// 1833 https://leetcode-cn.com/problems/maximum-ice-cream-bars/
// 先排序 然后贪心算法
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	less := coins
	sum := 0
	for len(costs) > 0 {
		cost := costs[0]
		costs = costs[1:]

		if cost > less {
			break
		}

		sum++
		less -= cost
	}

	return sum
}

// 另一种 贪心解法 时间更优
func maxIceCream2(costs []int, coins int) int {
	m := [1e5 + 1]int{}

	// // 列出花费为 cost 的棒冰有多少只
	for _, cost := range costs {
		m[cost]++
	}

	sum := 0
	// 贪心算法 从小开始取
	for i := 1; i <= 1e5 && coins >= i; i++ {
		if m[i] == 0 {
			continue
		}

		// 买下全部花费为 i 的棒冰或买下最多能买的数量
		can := min(m[i], coins/i)
		coins -= can * i

		sum += can
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
