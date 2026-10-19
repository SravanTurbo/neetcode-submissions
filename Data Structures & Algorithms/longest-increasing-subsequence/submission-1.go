/*
start: 7:40
idea: dp
t: O(n*n) s: O(n)
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)

	globalMax := 0

	dp[0] = 0
	for i:=1; i<n+1; i++ {
		maxi := 0
		for j:=0; j<i; j++ {
			if j == 0 || nums[j-1] < nums[i-1] {
				maxi = max(maxi, 1 + dp[j])
			}
		}
		dp[i] = maxi
		globalMax = max(globalMax, dp[i])
	}

	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
