package main

// 789 https://leetcode-cn.com/problems/escape-the-ghosts/
// 如果中途能碰到鬼，那么鬼一定可以同时或者先比我们到达该点
// 而如果存在这样的情况，那鬼一定可以同时或者比我们先到达终点
func escapeGhosts(ghosts [][]int, target []int) bool {
	start := []int{0, 0}
	dis := manDis(start, target)
	for _, g := range ghosts {
		if manDis(g, target) <= dis {
			return false
		}
	}
	return true
}

func manDis(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
