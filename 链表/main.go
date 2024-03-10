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

/*
func isPalindrome(head *ListNode) bool {
	if head == nil {
        return true
    }

    firstHeadEnd := endofFirstHalf(head)
    secondHalfStart := reverseList(firstHeadEnd.Next)

    p1 := head
    p2 := secondHalfStart
    result := true
    for result && p2 != nil {
        if p1.Val != p2.Val {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    firstHeadEnd.Next = reverseList(secondHalfStart)
    return result
}

func reverseList(head *ListNode) *ListNode {
    var prev,cur *ListNode = nil,head
    for cur != nil {
        temp := cur.Next
        cur.Next = prev
        prev = cur
        cur = temp
    }
    return prev
}

func endofFirstHalf (head *ListNode) *ListNode {
     fast := head
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
*/

// 21.环形链表
func hasCycle(head *ListNode) bool {
	var headMap = map[*ListNode]int{}
	h := head
	for h != nil {
		if _, ok := headMap[h]; ok {
			return true
		}
		headMap[h] = 1
		h = h.Next
	}
	return false
}

// 22环形链表 II
func detectCycle(head *ListNode) *ListNode {
	var headMap = map[*ListNode]int{}
	p := head
	i := 0
	for p != nil {
		if _, ok := headMap[p]; ok {
			return p
		}
		headMap[p] = i
		i++
		p = p.Next
	}
	return nil
}

// 23.合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 特殊情况
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p1, p2 := list1, list2
	var pHead *ListNode = nil
	//首先拿到头
	if p1.Val < p2.Val {
		pHead = p1
		p1 = p1.Next
	} else {
		pHead = p2
		p2 = p2.Next
	}
	temp := pHead
	// 循环
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			temp.Next = p1
			p1 = p1.Next
			temp = temp.Next
		} else {
			temp.Next = p2
			p2 = p2.Next
			temp = temp.Next
		}
	}
	if p1 == nil {
		temp.Next = p2
	} else {
		temp.Next = p1
	}
	return pHead
}

// TODO: 24-->LRU

type LRUCache struct {
	curIndex int
	capacity int
	valuesKV map[int]int
	valuesVK map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		curIndex: 0,
		capacity: capacity,
		valuesKV: make(map[int]int, capacity),
		valuesVK: make(map[int]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.valuesKV[key]; ok {
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.curIndex >= this.capacity {
		// 删除
		dKey := this.valuesVK[this.curIndex-this.capacity]
		delete(this.valuesVK, this.curIndex-this.capacity)
		delete(this.valuesKV, dKey)
		// 添加
		this.valuesKV[key] = value
		this.valuesVK[this.curIndex] = key
		this.curIndex++
	} else {
		// 存入
		this.valuesKV[key] = value
		this.valuesVK[this.curIndex] = key
		this.curIndex++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// 25-->两数相加
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	dump := &ListNode{}
	cur := dump
	for l1 != nil || l2 != nil || add != 0 {
		val := add
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: val % 10}
		add = val / 10
		cur = cur.Next
	}
	return dump.Next
}

// 26. 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 27 .24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummy.Next
}

// TODO： 28. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	}
	if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			next = cur
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			cur = next
			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}
	return dummyHead.Next
}
