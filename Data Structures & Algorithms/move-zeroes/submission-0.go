/*
start: 1:07
idea: 2-pointer
t: O(n) s:O(1)
*/
func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) && l < len(nums) {
		if nums[l] != 0 {
			l++
		} else if r < l {
			r = l+1
		} else if nums[r] == 0 {
			r++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
		}
 	}
}

