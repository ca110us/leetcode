package main

import (
	"sort"
	"strings"
)

// 面试 10.02 https://leetcode-cn.com/problems/group-anagrams-lcci/
// 把每个字符串排序 存入 map 最后遍历 map
// 如 ate eat tea 的 key 都为 aet
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)

		sortS := strings.Join(s, "")
		m[sortS] = append(m[sortS], str)
	}

	ans := make([][]string, 0)
	for _, strs := range m {
		a := []string{}
		for _, str := range strs {
			a = append(a, str)
		}
		ans = append(ans, a)
	}

	return ans
}
