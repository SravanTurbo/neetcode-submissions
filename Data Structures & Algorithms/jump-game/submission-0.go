/*
start: 12:26
idea: greedy - trversing from dest to start
t: O(n) s: O(1)
*/
func canJump(nums []int) bool {
	dest := len(nums)-1
	prev := len(nums)-2

	for prev >= 0 {
		if nums[prev] >= dest - prev {
			dest = prev
		}
		prev--
	}

	return dest == 0
}
