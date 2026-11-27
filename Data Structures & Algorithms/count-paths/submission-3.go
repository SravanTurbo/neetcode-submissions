/*
start: 12:17
idea: dp
t:O(m*n); s:O(m*n)
*/
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	memo := make(map[string]int)

	var numPaths func(r, c int) int
	numPaths = func(r,c int) int {
		if r == 1 || c == 1 {
			return 1
		}

		key := fmt.Sprintf("%d,%d", r, c)
		val, ok := memo[key]
		if ok {
			return val
		}

		memo[key] = numPaths(r-1, c) + numPaths(r, c-1)
		return memo[key]
	}

	return numPaths(m, n)
}
