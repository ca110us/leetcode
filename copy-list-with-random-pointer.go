package main

// 138 https://leetcode-cn.com/problems/copy-list-with-random-pointer/
// 哈希 哈希的 key 是旧链表，有指向关系 哈希的 value 是新链表 只是结构
func copyRandomList(head *Node) *Node {
	h := map[*Node]*Node{}
	for p := head; p != nil; p = p.Next {
		h[p] = &Node{p.Val, nil, nil}
	}
	for p := head; p != nil; p = p.Next {
		h[p].Next = h[p.Next]
		h[p].Random = h[p.Random]
	}
	return h[head]
}
