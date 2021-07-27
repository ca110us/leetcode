package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 671 https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
// 深度优先搜索 dfs 搜索结点值和根结点值一样的子树
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}

	if root.Left.Val == root.Right.Val && root.Left.Left == nil {
		return -1
	}

	sm := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Val == root.Val {
			if node.Left == nil {
				return
			}

			dfs(node.Left)
			dfs(node.Right)
		} else {
			if sm == -1 {
				sm = node.Val
			} else {
				sm = min(sm, node.Val)
			}
		}
	}

	dfs(root)
	return sm
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
