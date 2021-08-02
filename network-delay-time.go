package main

import (
	"math"
)

// 743 https://leetcode-cn.com/problems/network-delay-time/
// Dijkstra 算法
func networkDelayTime(times [][]int, n int, k int) int {
	// 点与点之间的距离，有向，不能过去的为无穷大
	const inf = math.MaxInt64 / 2
	grath := make([][]int, n)
	//  赋初始值的原因是因为点与点之间有向，无法过去的距离为无穷大
	for i := range grath {
		grath[i] = make([]int, n)
		for j := range grath[i] {
			grath[i][j] = inf
		}
	}
	for i := range times {
		//  将节点往前移动一位
		row, col := times[i][0]-1, times[i][1]-1
		grath[row][col] = times[i][2]
	}
	//  当前点距离k点的最小距离
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = inf
	}
	// 寻找当前点与k点距离的最小
	distance[k-1] = 0
	// used表示这个节点是否被使用，只有k点有最小距离才会被使用
	used := make([]bool, n)
	for i := 1; i < n; i++ {
		// 距离k点最小的值的节点的坐标
		x := -1
		// 寻找距离k点距离最小的节点的坐标
		for j, u := range used {
			// u为ture是表示这个节点被使用，要往后推一个
			// x == -1只是为了取到第一个点
			// distance[j] < distance[x] 目的是寻找最小点
			// 第一次就是要找到k点，之后就是在k点基础上寻找最小距离的点
			// 并且，最小路径找到后，!u可以将在k点以后未参与到最小路径的节点找出来
			// 将其加到对应的路径上，构建完整的图
			if !u && (x == -1 || distance[j] < distance[x]) {
				x = j
			}
		}
		used[x] = true
		// 距离k点距离最小的点找到了，现在开始取这个点到达所有目的节点的距离
		for i, dis := range grath[x] {
			// grath[x]中的点并不是所有都是其的目的节点，
			// 只有是x的目的节点才参与比较，不是的话，会被保留其最小值
			// 因为参与相加的还有其距离，如果不是x的目的节点，它的距离很大
			// 这样会保留原本的值
			distance[i] = min(distance[i], dis+distance[x])
		}
	}
	ans := 0
	for _, d := range distance {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return ans
}
func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
