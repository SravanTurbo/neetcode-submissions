/*
start: 7:53
t: O(n), s:O(n)
*/
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i:=1; i<len(nums); i++ {
		dp[i+1] = max(dp[i], dp[i-1] + nums[i])
	}
    
	return dp[len(nums)]
}
