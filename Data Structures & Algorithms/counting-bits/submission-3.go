func countBits(n int) []int {
	dp := make([]int, n+1)
	offset := 1

	for i:=1; i<n+1; i++ {
		if i/2 >= offset {
			offset = i
		}

		dp[i] = 1 + dp[i-offset]
	}
	
	return dp
}
