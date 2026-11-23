/*
start: 2:00pm
idea: marking visited
t: O(n) s: O(1)
end: 2:14pm
*/
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	i := 0
	for i >= 0 {
		if nums[i] == 0 {
			return i
		}

		temp := nums[i]
		nums[i] = 0 //visited
		i = temp
	}

	return -1
}
