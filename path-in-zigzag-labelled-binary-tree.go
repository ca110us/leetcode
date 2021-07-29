package main

import "sort"

// 1104 https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/
// 用满二叉树的公式 最左边节点 2^(i-1) 最右边节点是 2^i-1
// 反过来的时候就是下面的函数 2^(i-1) + 2^i-1 - label(从左到右排的话是第几个)
func reverse(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}

func pathInZigZagTree(label int) (path []int) {
	row := 1
	for 1<<row-1 < label {
		row++
	}
	path = append(path, label)

	if row&1 == 0 {
		label = reverse(label, row)
	}

	for row > 1 {
		label >>= 1
		row--
		if row&1 == 0 {
			path = append(path, reverse(label, row))
		} else {
			path = append(path, label)
		}
	}

	sort.Ints(path)
	return
}
