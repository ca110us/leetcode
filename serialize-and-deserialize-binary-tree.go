package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 297 https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/submissions/
// 这里使用 BFS 递归，也可以用 DFS + 队列去做
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func toTree(list *[]string) *TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: val}
	root.Left = toTree(list)
	root.Right = toTree(list)
	return root
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}

	str := fmt.Sprintf("%s,%s,%s", strconv.Itoa(root.Val), this.serialize(root.Left), this.serialize(root.Right))
	return str
}

func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	return toTree(&list)
}
