package main

// 10.
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。
func subarraySum(nums []int, k int) int {
	ans := 0
	pre := 0
	mp := map[int]int{}
	mp[0] = 1
	for _, v := range nums {
		pre += v
		if _, ok := mp[pre-k]; ok {
			ans += mp[pre-k]
		}
		mp[pre]++
	}
	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for i < k-1 {
			for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
				queue = queue[:len(queue)-1]
			}
			queue = append(queue, i)
			i++
		}
		// i == k-1 加到最后的k位置
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 取最大元素
		ans = append(ans, nums[queue[0]])
		// 当前最大值的下标到达了 ，弹出
		if i-queue[0]+1 == k {
			queue = queue[1:]
		}
	}
	return ans
}

// 76. 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	left, right, valid := 0, 0, 0
	start := 0
	l := len(s) + 1
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		// 移动窗口
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < l {
				start = left
				l = right - left
			}
			// d 是将移出窗口的字符
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if l == len(s)+1 {
		return ""
	} else {
		return s[start : start+l]
	}
}
