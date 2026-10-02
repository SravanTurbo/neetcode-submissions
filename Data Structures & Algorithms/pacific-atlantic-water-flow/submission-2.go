func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])

	output := [][]int{}

	pacific  := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for r := range rows {
		pacific[r]  = make([]bool, cols)
		atlantic[r] = make([]bool, cols)
	}

	var dfs func(r, c, prev int, valid [][]bool) 
	dfs = func(r, c, prev int, valid [][]bool) {
		if r < 0 || c < 0 || c >= cols || r >= rows {
			return
		}

		if heights[r][c] < prev {
			return
		}

		valid[r][c] = true

		tmp := heights[r][c]
		heights[r][c] = -1

		dfs(r+1, c, tmp, valid)
		dfs(r-1, c, tmp, valid)
		dfs(r, c+1, tmp, valid)
		dfs(r, c-1, tmp, valid)

		heights[r][c] = tmp
	}

	for r := range rows {
		for c := range cols {
			if r == 0 || c == 0 {
				dfs(r, c, -1, pacific)
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if r == rows-1 || c == cols-1 {
				dfs(r, c, -1 , atlantic)
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if pacific[r][c] && atlantic[r][c] {
				output = append(output, []int{r,c})
			}
		}
	}
	
	return output
}
