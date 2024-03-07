package main

import (
	"fmt"
	"math"
	"sort"
)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	res := 0
	l1, l2 := 1, 2
	for i := 3; i <= n; i++ {
		res = l1 + l2
		l1 = l2
		l2 = res
	}
	return res
}

func generate(numRows int) [][]int {
	// dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}

func rob(nums []int) int {
	// dp[i] = max(dp[i-1],dp[i-2] + nums[i],dp[i-3]+nums[i])
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	if n == 3 {
		return max(nums[0], nums[1], nums[2]+nums[0])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	dp[2] = max(nums[1], nums[0]+nums[2])
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i], dp[i-3]+nums[i])
	}
	return dp[n-1]
}

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; i >= j*j; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

func coinChange(coins []int, amount int) int {

	// dp[i] = 1 + min(dp[i-x])
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		minn := math.MaxInt32
		for j := 0; j < len(coins) && i >= coins[j]; j++ {
			minn = min(minn, dp[i-coins[j]])
		}
		dp[i] = minn + 1
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j]
	// init dp
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = 1
	// loop update  dp[i][j] = dp[i-1][j] + dp[i][j]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j]+dp[i][j-1], dp[i][j])
		}
	}
	return dp[m][n]
}

func minPathSum(grid [][]int) int {
	// init dp
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// dp[i][j] = min(dp[i-1][j] + grid[i-1][j-1],dp[i][j-1]+grid[i-1][j-1])
	// init side value
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(1, 1))
}
