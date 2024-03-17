package main

import (
	"fmt"
	"sort"
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

func to62RadixString(seq int64) string {
	ch := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
		'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	ans := make([]byte, 0)

	for {
		remainder := int(seq % 62)
		ans = append(ans, ch[remainder])
		seq /= 62
		if seq == 0 {
			break
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return i > j
	})
	return string(ans)
}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])

	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	ans := false
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		// 边界检查
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || vis[i][j] || board[i][j] != word[k] || ans {
			return
		}

		if k == len(word)-1 {
			ans = true
			return
		}
		vis[i][j] = true
		dfs(i+1, j, k+1)
		dfs(i-1, j, k+1)
		dfs(i, j+1, k+1)
		dfs(i, j-1, k+1)
		vis[i][j] = false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dfs(i, j, 0)
		}
	}
	return ans
}

func main() {
	n := 3
	println(n / 2)

}
