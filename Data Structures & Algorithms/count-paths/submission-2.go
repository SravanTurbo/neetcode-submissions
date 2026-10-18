/*
idea: recursion
t: O(m*n), s:O(m*n)
*/
func uniquePaths(m int, n int) int {
	memo := make(map[string]int)
	return uniquePathHelper(m, n, memo)
}

func uniquePathHelper(m, n int, memo map[string]int) int {
	if m == 1 || n == 1 {
		return 1
	}

	key := fmt.Sprintf("%d,%d", m, n)
	val, ok := memo[key]
	if ok {
		return val
	}

	memo[key] = uniquePathHelper(m-1, n, memo) + uniquePathHelper(m, n-1, memo)
	return memo[key]
}


