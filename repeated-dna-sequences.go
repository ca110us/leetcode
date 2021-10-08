package main

// 187 https://leetcode-cn.com/problems/repeated-dna-sequences/
// 哈希表
func findRepeatedDnaSequences(s string) (res []string) {
	m := map[string]int{}
	l := len(s)
	for i := 0; i <= l-10; i++ {
		m[s[i:i+10]]++
	}

	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return
}
