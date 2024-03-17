package main

// 35. 搜索插入位置

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	m, n := len(matrix), len(matrix[0])
	for row < m && matrix[row][n-1] < target {
		row++
	}
	if row == m {
		return false
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if target == matrix[row][mid] {
			return true
		}
		if target > matrix[row][mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
//func searchRange(nums []int, target int) []int {
//
//}
