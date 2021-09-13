package main

// 447 https://leetcode-cn.com/problems/number-of-boomerangs/
// 模拟 + hash
func numberOfBoomerangs(points [][]int) (ans int) {
	for _, p := range points {
		m := map[int]int{}
		for _, check := range points {
			dis := (p[0]-check[0])*(p[0]-check[0]) + (p[1]-check[1])*(p[1]-check[1])
			m[dis]++
		}
		for _, cnt := range m {
			ans += cnt * (cnt - 1)
		}
	}
	return
}
