package mian

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
