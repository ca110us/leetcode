package main

// 1588 https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/
// 前缀和
func sumOddLengthSubarrays(arr []int) (ans int) {
	ln := len(arr)
	for i := 1; i < ln; i++ {
		arr[i] += arr[i-1]
	}
	for i := 1; i <= ln; i += 2 {
		a, b := 0, i
		ans += arr[b-1]
		for b < ln {
			ans += arr[b] - arr[a]
			a++
			b++
		}
	}
	return
}
