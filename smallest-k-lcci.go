package main

import "sort"

// https://leetcode-cn.com/problems/smallest-k-lcci/
func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
