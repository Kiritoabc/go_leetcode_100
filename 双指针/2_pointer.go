package main

// 4.
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func moveZeroes(nums []int) {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			if l != r {
				nums[l] = nums[r]
				nums[r] = 0
			}
			l++
		}
	}
}
