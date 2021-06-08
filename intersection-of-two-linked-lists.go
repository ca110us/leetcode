package main

// 160 https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// 头尾相连，遍历完的时候要么指到一起，要么都是 nil
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	if a == nil || b == nil {
		return nil
	}

	for {
		if a != b {
			if a != nil {
				a = a.Next
			} else {
				a = headB
			}

			if b != nil {
				b = b.Next
			} else {
				b = headA
			}
		} else {
			break
		}
	}

	return a
}
