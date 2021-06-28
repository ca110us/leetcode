package main

// 815 https://leetcode-cn.com/problems/bus-routes/
// BFS 看注释吧 和 752 909 大同小异
func numBusesToDestination(routes [][]int, source int, target int) int {
	// 当前就在了 不用换乘啦
	if source == target {
		return 0
	}
	// 记录站点能进入的路线
	s2r := make(map[int][]int)
	// 经过的路线
	q := []int{}
	// 进入路线换乘的次数
	countMap := make(map[int]int)

	for i, r := range routes {
		for _, s := range r {
			if s == source {
				q = append(q, i)
				countMap[i] = 1
			}

			if _, ok := s2r[s]; ok {
				s2r[s] = add(s2r[s], i)
			} else {
				rec := []int{i}
				s2r[s] = rec
			}
		}
	}

	for len(q) > 0 {
		now := q[0]
		q = q[1:]

		step := countMap[now]
		for _, s := range routes[now] {
			if s == target {
				return step
			}

			// 将新路线加入队列
			if rs, ok := s2r[s]; ok {
				for _, r := range rs {
					if _, ok := countMap[r]; !ok {
						q = append(q, r)
						countMap[r] = step + 1
					}
				}
			}
		}
	}

	return -1
}

// 保证切片元素唯一
func add(s []int, e int) []int {
	for _, v := range s {
		if v == e {
			return s
		}
	}

	s = append(s, e)
	return s
}
