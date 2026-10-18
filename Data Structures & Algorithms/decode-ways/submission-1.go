/*
start: 9:50
idea: dp

*/
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		char := s[i-1]
		if char != '0' {
			dp[i] += dp[i-1]
		}

		val, _ := strconv.Atoi(s[i-2:i])
		if val >= 10 && val <=26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
