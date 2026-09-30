func exist(board [][]byte, word string) bool {	
	var dfs func(int, int, int, [][]bool) bool 
	
	dfs = func(i, j, k int, visited [][]bool) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		k++
		visited[i][j] = true

		directions := [][2]int{{1,0},{0,1},{-1,0},{0,-1}}
		for _, dir := range directions {
			x := i + dir[0]
			y := j + dir[1]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && !visited[x][y] {
				if dfs(x, y, k, visited) {
					return true
				}
			}
		}

		visited[i][j] = false

		return false			
	}

	visited := make([][]bool, len(board))
	for p := range len(board) {
		visited[p] = make([]bool, len(board[0]))
	}
	
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}

