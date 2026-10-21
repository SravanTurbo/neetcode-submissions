/*
start: 1:37
idea: 2-pointer
t: O(n) s: O(1)
end: 1:45
*/
func sortColors(nums []int) {
	l := 0
	for r:=0; r<len(nums); r++ {
		if nums[r] == 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}

	for r:=l; r<len(nums); r++ {
		if nums[r] == 1 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
