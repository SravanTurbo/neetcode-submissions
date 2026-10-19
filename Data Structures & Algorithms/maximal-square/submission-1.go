/*
idea: dp bottom-up
t: O(m*n) s: O(m*n)
*/
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows)
	for r := range rows {
		dp[r] = make([]int, cols)
	}

	maxLength := 0
	for r := range rows {
		for c := range cols {
			if r == 0 || c == 0 || matrix[r][c] == '0' {
				dp[r][c], _ = strconv.Atoi(string(matrix[r][c]))
			} else {
				dp[r][c] = 1 + min(dp[r-1][c], dp[r][c-1], dp[r-1][c-1])
			}

			maxLength = max(maxLength, dp[r][c])
		}
	}
	return maxLength*maxLength	
}

func min(a, b, c int) int {
	if a <= b && a <= c{
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
