package main

// 2 https://leetcode-cn.com/problems/add-two-numbers/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := new(ListNode)
	curr := node

	less := 0
	for l1 != nil || l2 != nil || less != 0 {
		curr.Next = new(ListNode)
		curr = curr.Next

		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Val = (sum + less) % 10
		less = (sum + less) / 10
	}
	return node.Next
}
