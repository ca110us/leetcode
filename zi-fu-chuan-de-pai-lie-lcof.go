package main

// 剑指 offer - 38 https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
// 回溯 + hash
// 可参考 https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/mian-shi-ti-38-zi-fu-chuan-de-pai-lie-hui-su-fa-by/
func permutation(s string) []string {
	res := make([]string, 0)
	c := []byte(s)

	var dfs func(x int)
	dfs = func(x int) {
		if x == len(s)-1 {
			res = append(res, string(c))
			return
		}

		dict := make(map[byte]int)
		for i := x; i < len(s); i++ {
			if _, ok := dict[c[i]]; ok {
				continue
			}
			dict[c[i]] = 1

			c[x], c[i] = c[i], c[x]
			// dfs(0)
			// a abc
			// dfs(1)
			// b bac
			// dfs(1)
			// c cba
			dfs(x + 1)
			c[x], c[i] = c[i], c[x]
		}
	}
	dfs(0)

	return res
}
