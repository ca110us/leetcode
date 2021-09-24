package main

// 430 https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/
// dfs
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	s := []*Node{root}
	help := &Node{Next: root}
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]
		node.Prev = help
		help.Next = node
		for node != nil {
			if node.Child != nil {
				if node.Next != nil {
					s = append(s, node.Next)
				}
				node.Next = node.Child
				node.Child.Prev = node
				node.Child = nil
			}
			help = node
			node = node.Next
		}
	}
	root.Prev = nil
	return root
}
