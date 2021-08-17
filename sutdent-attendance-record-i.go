package main

// 551 https://leetcode-cn.com/problems/student-attendance-record-i/
// 条件判断
func checkRecord(s string) bool {
	lNum := 0
	aCnt := 0

	for _, str := range s {
		if str == 'A' {
			aCnt++
			lNum = 0
		}
		if str == 'L' {
			lNum++
		}
		if str == 'P' {
			lNum = 0
		}

		if aCnt >= 2 || lNum >= 3 {
			return false
		}
	}
	return true
}
