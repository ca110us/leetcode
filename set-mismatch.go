package main

// 645 https://leetcode-cn.com/problems/set-mismatch/
// 计算出错误总和，然后通过重复值、错误总和、正确总和计算出丢失值
func findErrorNums(nums []int) []int {
	ln := len(nums)
	m := make([]int, ln+1)
	wrongSum := 0
	repeatNum := 0

	for _, num := range nums {
		if repeatNum == 0 {
			m[num]++
			if m[num] > 1 {
				repeatNum = num
			}
		}

		wrongSum += num
	}

	correctSum := (1 + ln) * ln / 2
	loseNum := repeatNum - (wrongSum - correctSum)

	return []int{repeatNum, loseNum}
}
