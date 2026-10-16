/*
start: 7:20pm

*/
func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0
	if n == 0 {
		return dp
	}
	
	dp[1] = 1
	if n == 1 {
		return dp
	}

	last := 1 //last perfectPower
	lastPower := 0 //lastPower of 2
	for i := 2; i < n+1; i++ {
		if (i%2) == 0 && i == int(math.Pow(2, float64(lastPower+1))){
			dp[i] = 1
			last = i
			lastPower += 1
			continue
		}

		dp[i] = dp[last] + dp[i-last]
 	}

	return dp
}
