package main

import "sort"

// 136. 只出现一次的数字

func singleNumber(nums []int) int {
	reduce := 0
	for i := 0; i < len(nums); i++ {
		reduce = reduce ^ nums[i]
	}

	return reduce
}

// 169. 多数元素

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 75. 颜色分类

func sortColors(nums []int) {
	a, b := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		for ; i <= b && nums[i] == 2; b-- {
			nums[i], nums[b] = nums[b], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[a] = nums[a], nums[i]
			a++
		}
	}
}

// 31. 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	// get
	var tail = n - 1
	for ; tail > 0; tail-- {
		if nums[tail-1] < nums[tail] {
			break
		}
	}
	for i := n - 1; i >= tail && tail > 0; i-- {
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

// 287. 寻找重复数

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {

	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
