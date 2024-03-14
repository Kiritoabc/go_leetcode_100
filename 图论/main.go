package main

// 200. 岛屿数量

func numIslands(grid [][]byte) int {
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x > len(grid)-1 || y > len(grid[0])-1 || y < 0 || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
