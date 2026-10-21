/*
start: 7:20
idea: 2 pointer
t:O(n) s:O(n)
end: 7:36
*/

func maximumSubarraySum(nums []int, k int) int64 {
	if len(nums) == 0 || len(nums) < k {
		return 0
	}

	maxScore := 0
	state := make(map[int]int)
	
	score := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		score += nums[end]
		state[nums[end]] += 1

		if end-start+1 == k {
			if len(state) == k {
				maxScore = max(maxScore, score)
			}

			score -= nums[start]
			state[nums[start]] -= 1
			if state[nums[start]] == 0 {
				delete(state, nums[start])
			}
			start += 1
		}
	}

	return int64(maxScore)
}
