package mian

import "sort"

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
