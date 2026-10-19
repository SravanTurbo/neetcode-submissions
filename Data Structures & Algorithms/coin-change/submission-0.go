/*
start: 10:07
*/
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		mini := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			coin := coins[j]
			if i < coin || dp[i-coin] == -1 {
				if dp[i] == 0 {
					dp[i] = -1
				}
				continue
			}

			change := 1 + dp[i-coin]
			mini    = min(change, mini)
			dp[i]   = mini
		}
	}

	return dp[amount]
}
