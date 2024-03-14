package main

import (
	"fmt"
	"sync"
)

func Loop(ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("send %d\n", v)
			wg.Done()
		}
	}
}

var wg sync.WaitGroup

var (
	n      = 1010
	parent []int
)

func initialize() {
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}

func find(u int) int {
	if u == parent[u] {
		return u
	}
	parent[u] = find(parent[u])
	return parent[u]
}

func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	parent[v] = u
}

func isSame(u, v int) bool {
	return find(u) == find(v)
}

func isValid(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}
	str := []byte{'(', ')', '{', '}', '[', ']'}
	stack := make([]byte, 0)
	//stack = append(stack, s[0])
	for i := 0; i < n; i++ {
		if s[i] == str[0] || s[i] == str[2] || s[i] == str[4] {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			ch := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for j := 0; j < 3; j++ {
				if s[i] == str[j*2+1] && ch != str[j*2] {
					return false
				}
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func nextPermutation(nums []int) {
	n := len(nums)
	// get
	var tail = n - 1
	for ; tail > 0; tail-- {
		if nums[tail-1] < nums[tail] {
			break
		}
	}
	if tail == 0 {
		l, r := tail, n-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
		return
	}
	for i := n - 1; i >= tail; i-- {
		if nums[i] > nums[tail-1] {
			nums[tail-1], nums[i] = nums[i], nums[tail-1]
			break
		}
	}
	l, r := tail, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	nextPermutation([]int{1, 3, 2})
}
