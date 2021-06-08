package main

// 203 https://leetcode-cn.com/problems/remove-linked-list-elements/
// 记录前一个节点，如果下一个 next 等于 val， 就把下一个的节点的 next 赋给 pre 的 next
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var h *ListNode
	var pre *ListNode

	if head == nil {
		return head
	}

	now := head
	for {
		if now.Val != val {
			if h == nil {
				h = now
			}

			pre = now
		} else {
			if pre != nil {
				pre.Next = now.Next
			}
		}

		if now.Next != nil {
			now = now.Next
		} else {
			break
		}
	}

	return h
}
