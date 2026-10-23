/*
start: 5:23pm
idea: dfs
t:O(n)  s: O(1)
end: 5:32
*/
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int) int 
	dfs = func (r, c int) int {
		if r<0 || c<0 || r>=rows || c>=cols {
			return 0
		}

		if grid[r][c] == 0 || grid[r][c] == -1 {
			return 0
		}

		grid[r][c] = -1 //marking visited

		return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
	} 

	area := 0
	maxArea := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] != 1 {
				continue
			}

			area = dfs(r,c)
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}
