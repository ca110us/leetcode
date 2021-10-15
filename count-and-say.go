package main

// 38 https://leetcode-cn.com/problems/count-and-say/
// 递归
func countAndSay(n int) string {
	return gen("1", n-1)
}

func gen(str string, n int) string {
	if n == 0 {
		return str
	}

	idx := 0
	res := []byte{}
	for idx < len(str) {
		c := str[idx]
		l := idx
		for j := idx; j < len(str); j++ {
			if str[j] != c {
				break
			}
			l = j
		}
		n := l - idx + 1
		res = append(res, byte(n)+'0')
		res = append(res, c)
		idx = l + 1
	}
	return gen(string(res), n-1)

}
