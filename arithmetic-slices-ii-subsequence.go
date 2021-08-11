package main

// 446 https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/
// 所有等差子序列问题
// 用 2 4 6 来解释大于等于两位的等差子序列问题
// 以 6 为尾项 差值为 2 的子序列数，等于以 4 为尾项 差值为 2 的子序列数 +1,
// 但是，以 6 为尾项的数量，要后面出现等差值元素后才能用
func numberOfArithmeticSlices(nums []int) (ans int) {
	m := make([]map[int]int, len(nums))

	for i, x := range nums {
		m[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := m[j][d]
			ans += cnt
			m[i][d] += cnt + 1
		}
	}
	return
}
