package main

// 10.
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	mp := map[int]int{}
	mp[0] = 1
	for i := range nums {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre]++
	}
	return count
}
