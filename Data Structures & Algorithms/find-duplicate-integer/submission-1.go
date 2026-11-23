/*
start: 2:00pm
idea: map
t: O(n) s: O(n)
end: 2:14pm
*/
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	unique := make(map[int]bool, len(nums))
	for i:=0; i<len(nums); i++ {
		elem := nums[i]
		_, exists := unique[elem]
		if exists {
			return elem
		}
		
		unique[elem] = true
	}

	return -1
}
