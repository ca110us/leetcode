package main

import (
	"container/heap"
	"math"
)

// 313 https://leetcode-cn.com/problems/super-ugly-number/
// 最小堆
func nthSuperUglyNumber(n int, primes []int) int {
	var ans int
	hash := make(map[int]bool)
	h := H{1}
	heap.Init(&h)

	for n > 0 {
		ans = heap.Pop(&h).(int)
		n--
		if n == 0 {
			return ans
		}
		for _, v := range primes {
			if !hash[ans*v] {
				heap.Push(&h, ans*v)
				hash[ans*v] = true
			}
		}
	}
	return -1
}

type H []int

func (h *H) Len() int {
	return len(*h)
}

func (h *H) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *H) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *H) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *H) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

// 动态规划
func nthSuperUglyNumber2(n int, primes []int) int {
	ln := len(primes)

	p := make([]int, ln)
	nums := make([]int, ln)
	for i := range p {
		p[i] = 1
	}
	f := make([]int, n+1)
	f[1] = 1

	for i := 2; i <= n; i++ {
		min := math.MaxInt32
		for j := range nums {
			nums[j] = f[p[j]] * primes[j]
			min = small(min, nums[j])
		}
		f[i] = min
		for j := range nums {
			if f[i] == nums[j] {
				p[j]++
			}
		}
	}
	return f[n]
}

func small(a, b int) int {
	if a < b {
		return a
	}
	return b
}
