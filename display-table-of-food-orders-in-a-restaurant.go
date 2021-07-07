package main

import (
	"sort"
	"strconv"
)

// 1418 https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant/
// Hash + 排序
func displayTable(orders [][]string) [][]string {
	var ans [][]string

	order := make(map[int]map[string]int)
	food := make(map[string]bool)
	for _, v := range orders {
		table, _ := strconv.Atoi(v[1])
		f := v[2]
		if order[table] == nil {
			order[table] = make(map[string]int)
		}
		order[table][f]++
		food[f] = true
	}

	var foods []string
	for k := range food {
		foods = append(foods, k)
	}
	sort.Strings(foods)

	title := []string{"Table"}
	for _, f := range foods {
		title = append(title, f)
	}

	ans = append(ans, title)
	var table []int
	for k := range order {
		table = append(table, k)
	}
	sort.Ints(table)

	for _, v := range table {
		var s []string
		id := strconv.Itoa(v)
		s = append(s, id)
		for _, b := range foods {
			num := strconv.Itoa(order[v][b])
			s = append(s, num)
		}
		ans = append(ans, s)
	}
	return ans
}
