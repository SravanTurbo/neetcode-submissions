func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs func(r, c, k int) bool 
	dfs = func(r, c, k int) bool {
		if k == len(word) {
			return true
		}

		if r<0 || c<0 || r>=rows || c>=cols ||
		   board[r][c] == '#' || board[r][c] != word[k] {
			return false
		}
		
		temp := board[r][c]
		board[r][c] = '#'
		res := dfs(r+1, c, k+1) ||
			   dfs(r-1, c, k+1) ||
			   dfs(r, c+1, k+1) ||
			   dfs(r, c-1, k+1)
		board[r][c] = temp

		return res
	}

	for i := range rows {
		for j := range cols {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
