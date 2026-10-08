/*
start: 9:30
t: O(n), s: O(n)
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) < 2 {
		return nums[0]
	}

	//rob from first to last-but-one: 0 to n-2
	dp1 := make([]int, len(nums))
	dp1[0] = 0
	dp1[1] = nums[0]
	for i:=1; i<len(nums)-1; i++ {
		dp1[i+1] = max(dp1[i], dp1[i-1] + nums[i])
	}

	//rob from second to last: 1 to n-1
	dp2 := make([]int, len(nums))
	dp2[0] = 0
	dp2[1] = nums[1]
	for i:=2; i<len(nums); i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2] + nums[i])
	}
	
	return max(dp1[len(nums)-1], dp2[len(nums)-1])
}
