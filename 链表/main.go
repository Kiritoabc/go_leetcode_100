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

// 19.反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func reverseList(head *ListNode) *ListNode {
	/*
		迭代：
		   var prev *ListNode
		    curr := head
		    for curr!=nil {
		        temp := curr.Next
		        curr.Next = prev
		        prev = curr
		        curr = temp
		    }
		    return prev
	*/

	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 20 回文链表
// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
func isPalindrome(head *ListNode) bool {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	n := len(values)
	for i, v := range values[:n/2] {
		if v != values[n-1-i] {
			return false
		}
	}

	return true
}
