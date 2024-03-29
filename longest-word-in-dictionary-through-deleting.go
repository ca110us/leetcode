package main

// 524 https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/
// 贪心
func findLongestWord(s string, dictionary []string) (ans string) {
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					if len(t) > len(ans) || len(t) == len(ans) && t < ans {
						ans = t
					}
					break
				}
			}
		}
	}
	return
}
