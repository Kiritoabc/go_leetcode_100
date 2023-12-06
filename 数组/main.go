package main

import "sort"

// 13.最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}

// 14.合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	// 首先根据左边排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		l1, r1 := intervals[i][0], intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			l2, r2 := intervals[j][0], intervals[j][1]
			// 右边比左边小，退出循环
			if r1 < l2 {
				break
			} else if r1 < r2 {
				// 否则合并
				r1 = r2
			}
			i++
		}
		// 添加进ans
		ans = append(ans, []int{l1, r1})
	}

	return ans
}

// 15.轮转数组
// 给定一个整数数组 nums，
// 将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// 16.除自身以外数组的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
func productExceptSelf(nums []int) []int {
	length := len(nums)
	ans := make([]int, len(nums))
	pre := make([]int, len(nums))
	tail := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			pre[i] = 1
			tail[length-i-1] = 1
		} else {
			pre[i] = nums[i-1] * pre[i-1]
			tail[length-i-1] = nums[length-i] * tail[length-i]
		}
	}

	for i := 0; i < length; i++ {
		ans[i] = pre[i] * tail[i]
	}
	return ans
}

// 17.缺失的第一个正数
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
