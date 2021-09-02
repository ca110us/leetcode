package main

// offer - 22 https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
// 滑动窗口法
func getKthFromEnd(head *ListNode, k int) *ListNode {
	m := make([]*ListNode, 0)
	node := head
	m = append(m, node)

	for node.Next != nil {
		if len(m) == k {
			m = m[1:]
		}
		node = node.Next
		m = append(m, node)
	}

	return m[0]
}

// 快慢指针法
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
