package main

// 541 https://leetcode-cn.com/problems/reverse-string-ii/
// 模拟 + 指针
func reverseStr(s string, k int) string {
	ln := len(s)
	already := 0

	str := []byte(s)
	for ln-already > 0 {
		reverseByIndex(&str, already, k)
		already += 2 * k
	}

	return string(str)
}

func reverseByIndex(s *[]byte, start, k int) {
	if start+k > len(*s) {
		k = len(*s) - start
	}

	idx := 0
	for idx <= (k-1)/2 {
		(*s)[start+idx], (*s)[start+k-1-idx] = (*s)[start+k-1-idx], (*s)[start+idx]
		idx++
	}
}
