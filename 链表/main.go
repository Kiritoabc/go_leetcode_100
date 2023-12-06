package main

// 18.相交链表
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := map[*ListNode]bool{}
	h1 := headA
	for h1 != nil {
		mp[h1] = true
		h1 = h1.Next
	}
	h2 := headB
	for h2 != nil {
		if mp[h2] {
			return h2
		}
		h2 = h2.Next
	}
	return nil
}
