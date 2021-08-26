package main

import "sort"

// 881 https://leetcode-cn.com/problems/boats-to-save-people/
// 贪心
func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)

	light, heavy := 0, len(people)-1
	for light <= heavy {
		if people[heavy] == limit {
			ans++
			heavy--
			continue
		}
		if people[light]+people[heavy] > limit {
			heavy--
		} else {
			heavy--
			light++
		}
		ans++
	}

	return
}
