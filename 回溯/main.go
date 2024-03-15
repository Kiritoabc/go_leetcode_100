package mian

import (
	"sort"
	"strings"
)

// 46. 全排列

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	st := make([]bool, len(nums))
	t := make([]int, 0, len(nums))

	var dfs func(cur int)

	dfs = func(cur int) {
		if cur == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, t)
			ans = append(ans, tmp)
		}

		for i := 0; i < len(nums); i++ {
			// in
			if !st[i] {
				t = append(t, nums[i])
				st[i] = true
				dfs(cur + 1)
				// 回溯
				st[i] = false
				t = t[:len(t)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// 78. 子集

func subsets(nums []int) [][]int {
	res, path := make([][]int, 0), make([]int, 0, len(nums))

	var dfs func(cur int)
	dfs = func(cur int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := cur; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

// 17. 电话号码的字母组合

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	ans := make([]string, 0)
	t := make([]byte, 0)
	nums := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var dfs func(cur int)
	// 参数
	dfs = func(cur int) {
		if len(t) == n {
			tmp := make([]byte, n)
			copy(tmp, t)
			ans = append(ans, string(tmp))
			return
		}
		str := nums[digits[cur]-'2']
		for i := 0; i < len(str); i++ {
			t = append(t, str[i])
			dfs(cur + 1)
			// 回溯
			t = t[:len(t)-1]
		}
	}
	dfs(0)
	return ans
}

//

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	t := make([]int, 0)
	n := len(candidates)
	sort.Ints(candidates)
	var dfs func(cur int, sum int)

	dfs = func(cur int, sum int) {
		if sum <= 0 {
			if sum == 0 {
				tmp := make([]int, len(t))
				copy(tmp, t)
				ans = append(ans, tmp)
			}
			return
		}
		//
		for i := cur; i < n; i++ {
			// 减
			if sum-candidates[i] < 0 {
				break
			}
			t = append(t, candidates[i])
			dfs(i, sum-candidates[i])
			t = t[:len(t)-1]
		}
	}

	dfs(0, target)
	return ans
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(left int, right int, s string)

	dfs = func(left int, right int, s string) {
		if left == n && right == n {
			ans = append(ans, s)
			return
		}
		if left < n {
			dfs(left+1, right, s+"(")
		}
		if right < left {
			dfs(left, right+1, s+")")
		}
	}
	dfs(0, 0, "")
	return ans
}

func partition(s string) [][]string {
	path := make([]string, 0)
	ans := make([][]string, 0)
	var dfs func(cur int)
	n := len(s)
	dfs = func(cur int) {
		if cur == n {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := cur; i < n; i++ {
			str := s[cur : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//79. 单词搜索

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])

	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	ans := false
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		// 边界检查
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || vis[i][j] || board[i][j] != word[k] || ans {
			return
		}

		if k == len(word)-1 {
			ans = true
			return
		}
		vis[i][j] = true
		dfs(i+1, j, k+1)
		dfs(i-1, j, k+1)
		dfs(i, j+1, k+1)
		dfs(i, j-1, k+1)
		vis[i][j] = false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dfs(i, j, 0)
		}
	}
	return ans
}

// 51. N 皇后

func solveNQueens(n int) [][]string {
	var res [][]string
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	var backtrack func(row int)
	backtrack = func(row int) {
		// get res
		if row == n {
			tmp := make([]string, n)
			for i, rowStr := range chessboard {
				tmp[i] = strings.Join(rowStr, "")
			}
			res = append(res, tmp)
		}

		// backtrack
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtrack(row + 1)
				chessboard[row][i] = "."
			}
		}

	}
	backtrack(0)
	return res
}

func isValid(n, row, col int, chessboard [][]string) bool {
	// 列 （只需要检查上面的列即可）
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 斜对角线 45 （只需要往上判断即可）
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	// 斜对角线 135 （只需要往上判断即可）
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
