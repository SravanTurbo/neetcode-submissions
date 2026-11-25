/*
start: 1:55pm
idea: local best at index
t: O(n) s: O(1)
end: 2:10pm
*/

func maxSubArray(nums []int) int {
	sum    := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		maxSum = max(maxSum, sum)
	}

	return maxSum
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}