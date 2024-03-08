package main

// 29  94. 二叉树的中序遍历

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func inorderTraversal(root *TreeNode) []int {
//	var res []int
//	var inorder func(node *TreeNode)
//	inorder = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		inorder(node.Left)
//		res = append(res, node.Val)
//		inorder(node.Right)
//	}
//	inorder(root)
//	return res
//}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 取出
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 30. 104. 二叉树的最大深度
//func maxDepth(root *TreeNode) int {
//	var res int = 0
//	var maxDepthRoot func(node *TreeNode, depth int) int
//	maxDepthRoot = func(node *TreeNode, depth int) int {
//		if node == nil {
//			return depth
//		}
//		depth += 1
//		l := maxDepthRoot(node.Left, depth)
//		r := maxDepthRoot(node.Right, depth)
//		if l > r {
//			return l
//		} else {
//			return r
//		}
//	}
//	res = maxDepthRoot(root, res)
//	return res
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 31. 226. 翻转二叉树

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// swap
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 31. 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var symmetric func(left *TreeNode, right *TreeNode) bool
	symmetric = func(left *TreeNode, right *TreeNode) bool {
		// false case
		if (left == nil && right != nil) || (right == nil && left != nil) {
			return false
		}

		if left != nil && right != nil && left.Val != right.Val {
			return false
		}

		// true case
		if left != nil && right != nil {
			return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
		}
		return true
	}

	return symmetric(root.Left, root.Right)
}

// 32. 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	var res = 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	depth(root)
	return res
}

// 33. 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func generateMatrix(n int) [][]int {
	nums := make([][]int, n)
	startX, startY := 0, 0
	offset := 1
	count := 1
	i, j := 0, 0
	for t := 0; t < n/2; t++ {
		for j = startY; j < n-offset; j++ {
			nums[i][j] = count
			count++
		}
		for i = startX; i < n-offset; i++ {
			nums[i][j] = count
			count++
		}
		for ; j > startY; j-- {
			nums[i][j] = count
			count++
		}
		for ; i > startX; i-- {
			nums[i][j] = count
			count++
		}
		startX++
		startY++
		offset++
	}

	if n%2 == 1 {
		nums[i][j] = count
	}
	return nums
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var dummyHead = &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
	}

	return dummyHead.Next
}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
	Pre  *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	dummyhead := this
	if dummyhead.Next == nil {
		return -1
	}
	for i := 0; i < index; i++ {
		dummyhead = dummyhead.Next
		if dummyhead.Next == nil {
			return -1
		}
	}
	return dummyhead.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := &MyLinkedList{Val: val}
	head.Pre = this
	head.Next = this.Next

	this.Next = head
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.Next == nil {
		cur = cur.Next
	}
	tail := &MyLinkedList{Val: val}
	// add tail
	cur.Next = tail
	tail.Next = nil
	tail.Pre = cur
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	dummyhead := this
	for i := 0; i < index; i++ {
		dummyhead = dummyhead.Next
		if dummyhead == nil {
			return
		}
	}
	insert := &MyLinkedList{Val: val}
	insert.Next = dummyhead.Next
	insert.Pre = dummyhead
	if dummyhead.Next != nil {
		dummyhead.Next.Pre = insert
	}
	dummyhead.Next = insert
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this
	if cur.Next == nil {
		return
	}
	for i := 0; i < index; i++ {
		cur = cur.Next
		if cur.Next == nil {
			return
		}
	}
	cur.Next = cur.Next.Next
	if cur.Next != nil {
		cur.Next.Pre = cur
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const (
	INT_MAX = int64(^uint64((0)) >> 1)
	INT_MIN = ^INT_MAX
)

func isValidBST(root *TreeNode) bool {
	return dfs(root, INT_MIN, INT_MAX)
}

func dfs(node *TreeNode, mmin, mmax int64) bool {
	if node == nil {
		return true
	}
	if int64(node.Val) <= mmin || int64(node.Val) >= mmax {
		return false
	}
	return dfs(node.Left, mmin, int64(node.Val)) && dfs(node.Right, int64(node.Val), mmax)
}
