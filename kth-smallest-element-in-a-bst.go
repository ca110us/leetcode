package main

// 230 https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 遍历
func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	find(root, &arr)
	return arr[k-1]

}
func find(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	find(root.Left, arr)
	*arr = append(*arr, root.Val)
	find(root.Right, arr)
	return

}
