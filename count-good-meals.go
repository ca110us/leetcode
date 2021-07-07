package main

// 1711 https://leetcode-cn.com/problems/count-good-meals/
// 哈希 + 位运算
func countPairs(deliciousness []int) int {
	m := make(map[int]int)
	ans := 0
	max := 0

	for _, d := range deliciousness {
		m[d]++
		if d > max {
			max = d
		}
	}

	for n1 := range m {
		for sum := 1; ; sum <<= 1 {
			n2 := sum - n1

			if n2 > max {
				break
			}

			if _, ok := m[n2]; !ok {
				continue
			}

			if n2 > n1 {
				ans += m[n1] * m[n2]
			} else if n2 == n1 {
				ans += m[n2] * (m[n2] - 1) / 2
			}
		}
	}

	return ans % (1e9 + 7)
}
