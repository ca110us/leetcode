package main

// 451 https://leetcode-cn.com/problems/sort-characters-by-frequency/
// 桶排序 从大往小取，可以令 max 直接为 len(s) 运行会变快，但是空间会变浪费
func frequencySort(s string) string {
	m := map[rune]int{}
	max := 0
	for _, v := range s {
		m[v]++
		max = bigger(m[v], max)
	}

	buckets := make([][]rune, max+1)
	for r, count := range m {
		buckets[count] = append(buckets[count], r)
	}

	runes := make([]rune, 0)
	for i := max; i >= 1; i-- {
		for _, r := range buckets[i] {
			runes = re(runes, r, i)
		}
	}

	return string(runes)
}

func re(arr []rune, r rune, count int) []rune {
	for i := 0; i < count; i++ {
		arr = append(arr, r)
	}

	return arr
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
