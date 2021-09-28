package main

// 437 https://leetcode-cn.com/problems/path-sum-iii/
// dfs
var mp = make(map[int]int)

func pathSum(root *TreeNode, targetSum int) int {
	mp[0] = 1
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, pre int, target int) int {
	if root == nil {
		return 0
	}
	cur := pre + root.Val
	res := mp[cur-target]
	mp[cur] = mp[cur] + 1
	left := dfs(root.Left, cur, target)
	right := dfs(root.Right, cur, target)
	mp[cur] = mp[cur] - 1
	return res + left + right
}
