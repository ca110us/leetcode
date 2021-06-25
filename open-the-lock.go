package main

import "strconv"

// 752 https://leetcode-cn.com/problems/open-the-lock/
// 广度优先搜索 BFS 双向搜索
func openLock(deadends []string, target string) int {
	// 不用算了
	if target == "0000" {
		return 0
	}

	deadMap := make(map[string]bool)
	for _, v := range deadends {
		deadMap[v] = true
	}

	// 还没开始就死了
	if _, ok := deadMap["0000"]; ok {
		return -1
	}

	// 从头和尾开始进行双向搜索
	q1 := []string{"0000"}
	q2 := []string{target}
	// 记录两个方向到这个 node 的步数
	visitMap1 := make(map[string]int)
	visitMap2 := make(map[string]int)
	// 头尾节点步数为 0
	visitMap1["0000"] = 0
	visitMap2[target] = 0

	t := -1
	for len(q1) != 0 && len(q2) != 0 {
		if len(q1) <= len(q2) {
			t, q1 = update(q1, visitMap1, visitMap2, deadMap)
		} else {
			t, q2 = update(q2, visitMap2, visitMap1, deadMap)
		}

		if t != -1 {
			return t
		}
	}
	return t
}

func update(queue []string, curVisit map[string]int, otherVisit map[string]int, deadMap map[string]bool) (int, []string) {
	node := queue[0]
	queue = queue[1:]

	step := curVisit[node]
	for i := 0; i < 4; i++ {
		for j := -1; j <= 1; j++ {
			if j == 0 {
				continue
			}
			num := int(node[i] - '0')

			n := 0
			// 该位向上拧得到的数字
			if j == 1 {
				n = (num + 1) % 10
			} else {
				n = (num + 9) % 10
			}

			pnode := node[:i] + strconv.Itoa(n) + node[i+1:]
			// 已经遍历到过了
			if _, ok := curVisit[pnode]; ok {
				continue
			}
			// 死亡节点 跳过
			if _, ok := deadMap[pnode]; ok {
				continue
			}

			// 找到了同一节点 就可以结束了
			if _, ok := otherVisit[pnode]; ok {
				return step + 1 + otherVisit[pnode], queue
			} else {
				queue = append(queue, pnode)
				curVisit[pnode] = step + 1
			}
		}
	}
	return -1, queue
}
