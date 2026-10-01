func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r<0 || r>=rows || c<0 || c>=cols {
			return
		}

		if grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0' //mark visited to water

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count 
}
