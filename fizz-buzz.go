package main

import "strconv"

// 412 https://leetcode-cn.com/problems/fizz-buzz/
// 位运算
func fizzBuzz(n int) []string {
	i := 1
	v := n + 1
	res := make([]string, n)
	for i < v {
		flag := 0
		if i%3 == 0 {
			flag = 1
		}
		if i%5 == 0 {
			flag = flag | 2
		}
		switch flag {
		case 1:
			res[i-1] = "Fizz"
		case 2:
			res[i-1] = "Buzz"
		case 3:
			res[i-1] = "FizzBuzz"
		default:
			res[i-1] = strconv.Itoa(i)
		}
		i++
	}
	return res
}
