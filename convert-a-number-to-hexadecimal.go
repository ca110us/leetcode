package main

// 405 https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/
// 进制转换
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	arr := make([]byte, 0, 8)
	for i := 7; i >= 0; i-- {
		n := byte(num >> (4 * i) & 0xf)
		if n > 0 || len(arr) > 0 {
			if n < 10 {
				arr = append(arr, n+'0')
			} else {
				arr = append(arr, n-10+'a')
			}
		}
	}
	return string(arr)
}
