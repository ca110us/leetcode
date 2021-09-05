package main

// 470 https://leetcode-cn.com/problems/implement-rand10-using-rand7/
// 拒绝采样
func rand10() int {
	for {
		r1 := rand7()
		r2 := rand7()
		num := r1 + (r2-1)*7
		if num <= 40 {
			return num%10 + 1
		}
	}
}
