/*
start: 10:13pm
idea: bfs from treasure chests - one step at a time from all chests
t: O(m*n) s: O(m*n) - queue
end: 11:13pm
*/

func islandsAndTreasure(grid [][]int) {
	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return 
	}

	queue := [][3]int{}
	visited := make([][]bool, rows)
	for r := range rows {
		visited[r] = make([]bool, cols)
		for c := range cols {
			if grid[r][c] == 0 {
				queue = append(queue, [3]int{r, c, 0})
			}
		}
	}

	for i := 0; i < len(queue); i++ {
		r, c, dist := queue[i][0], queue[i][1], queue[i][2]
		if r < 0 || c < 0 || r >= rows || c >= cols {
			continue
		}

		if visited[r][c] || grid[r][c] < dist {
			continue
		}

		grid[r][c]    = dist
		visited[r][c] = true
	
		queue = append(queue, [3]int{r+1, c, dist + 1})
		queue = append(queue, [3]int{r-1, c, dist + 1})
		queue = append(queue, [3]int{r, c+1, dist + 1})
		queue = append(queue, [3]int{r, c-1, dist + 1})
	}
}