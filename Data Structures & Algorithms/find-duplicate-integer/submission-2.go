/*
start: 2:15pm
idea: sorting
t: O(n log n) s: O(1)
end: 2:16pm
*/
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sort.Ints(nums)
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return -1
}
