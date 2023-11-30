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
