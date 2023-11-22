package main

import "sort"

// 1.
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if temp == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 2.
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// 3.
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
func longestConsecutive(nums []int) int {
	mp := map[int]int{}
	max_l := 0
	for i := range nums {
		num := nums[i]
		if mp[num] == 0 {
			// 左侧连续的数量
			left := mp[num-1]
			// 右侧连续的数量
			right := mp[num+1]
			cur_l := 1 + left + right
			if cur_l > max_l {
				max_l = cur_l
			}
			// 更新左右两侧
			mp[num] = cur_l
			mp[num-left] = cur_l
			mp[num+right] = cur_l
		}
	}
	return max_l
}
