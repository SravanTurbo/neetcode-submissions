/*
start: 12:17
idea: dp
t:O(m*n); s:O(m*n)
*/
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for r := range m {
		dp[r] = make([]int, n)
	}

	dp[0][0] = 1
	var numPaths func(r, c int) int
	numPaths = func(r,c int) int {
		if r < 0 || r >= m || c < 0 || c>=n {
			return 0
		}

		if dp[r][c] == -1 {
			return 0
		}

		if dp[r][c] != 0 {
			return dp[r][c]
		}

		dp[r][c] = numPaths(r-1, c) + numPaths(r, c-1)

		if dp[r][c] == 0 {
			dp[r][c] = -1 //mark visited
		}

		return dp[r][c]
	}

	numPaths(m-1, n-1)

	return dp[m-1][n-1]
}
