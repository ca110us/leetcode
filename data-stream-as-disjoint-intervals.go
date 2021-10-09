package main

import "sort"

// 352 https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals/
// 哈希表 + 并查集
type SummaryRanges struct {
	rangeMap map[int]*Range    //map of all values to the ranges
	rootMap  map[*Range]*Range //map of all root ranges
}

type Range struct {
	min, max int
	parent   *Range
	d        int
}

func newRange(val int) *Range {
	n := &Range{max: val, min: val}
	n.parent = n
	n.d = 1
	return n
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	sum := SummaryRanges{rangeMap: make(map[int]*Range), rootMap: make(map[*Range]*Range)}
	return sum
}

func (this *SummaryRanges) AddNum(val int) {
	if _, ok := this.rangeMap[val]; ok {
		return
	}
	// create new range
	n := newRange(val)
	this.rangeMap[val] = n
	this.rootMap[n] = n

	// check left of val
	if r, ok := this.rangeMap[val-1]; ok {
		this.union(r, n)
	}
	// check right of val
	r, ok := this.rangeMap[val+1]
	if !ok {
		return
	}
	// merge right
	this.union(n, r)
}

func (this *SummaryRanges) union(l, r *Range) {
	pl := find(l)
	pr := find(r)
	if pl.d > pr.d {
		pr.parent = pl
		pl.max = pr.max
		this.rootMap[pr] = nil
	} else {
		if pl.d == pr.d {
			pr.d++
		}
		pl.parent = pr
		pr.min = pl.min
		this.rootMap[pl] = nil
	}
}

func find(l *Range) *Range {
	if l.parent == l {
		return l
	}
	p := find(l.parent)
	l.parent = p
	return p
}

type Ranges []*Range

func (rs Ranges) Len() int {
	return len(rs)
}

func (rs Ranges) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs Ranges) Less(i, j int) bool {
	return rs[i].min < rs[j].min
}

func (this *SummaryRanges) GetIntervals() (results [][]int) {
	var rets Ranges
	for _, root := range this.rootMap {
		if root != nil {
			rets = append(rets, root)
		}
	}
	if len(rets) >= 2 {
		sort.Sort(rets)
	}
	for _, ret := range rets {
		results = append(results, []int{ret.min, ret.max})
	}
	return
}
