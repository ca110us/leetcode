package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 863 https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
// BFS 两次 第一次找出所有节点的父节点 第二次找出所有解
func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	p := map[int]*TreeNode{}
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			p[node.Left.Val] = node
			findParents(node.Left)
		}

		if node.Right != nil {
			p[node.Right.Val] = node
			findParents(node.Right)
		}
	}

	findParents(root)

	var findAns func(*TreeNode, *TreeNode, int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == k {
			ans = append(ans, node.Val)
			return
		}

		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if p[node.Val] != from {
			findAns(p[node.Val], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return
}
