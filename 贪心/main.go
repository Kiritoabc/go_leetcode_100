package main

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	mmin := prices[0]
	result := 0
	for i := 0; i < len(prices); i++ {
		if mmin > prices[i] {
			mmin = prices[i]
		}
		result = max(result, prices[i]-mmin)
	}
	return result
}

// 跳跃游戏
func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(cover, nums[i]+i)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 45. 跳跃游戏 II

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cover := 0
	tmp := 0
	ans := 0
	for i := 0; i <= cover; i++ {
		tmp = max(tmp, nums[i]+i)
		if i == cover {
			ans++
			cover = tmp
			if cover >= len(nums)-1 {
				return ans
			}
		}
	}
	return -1
}

// 763. 划分字母区间
func partitionLabels(s string) []int {
	hmap := make(map[byte]int)
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		hmap[s[i]] = i
	}
	maxP := hmap[0]
	startP := 0
	for i := 0; i < len(s); i++ {
		maxP = max(maxP, hmap[s[i]])
		if i == maxP {
			ans = append(ans, maxP-startP+1)
			startP = maxP + 1
		}
	}
	return ans
}
