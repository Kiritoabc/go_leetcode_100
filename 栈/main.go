package main

// 20. 有效的括号

func isValid(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}
	str := []byte{'(', ')', '{', '}', '[', ']'}
	stack := make([]byte, n)
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

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	queue := make([]int, 0)
	for i, v := range temperatures {
		for len(queue) != 0 && v > temperatures[queue[len(queue)-1]] {
			// 计算
			ans[queue[len(queue)-1]] = i - queue[len(queue)-1]
			// 弹出
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	// 没弹出的说明没有比他大的
	for len(queue) > 0 {
		ans[queue[0]] = 0
		queue = queue[1:]
	}
	return ans
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	mono_stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			right[i] = n
		} else {
			right[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

// 2,4

// left [-1,0]

// right [2,2]
