/*
start:12:16
idea: hashmap
*/
func twoSum(nums []int, target int) []int {
	firstIdx := make(map[int]int)

	for i:=0; i<len(nums); i++ {
		idx, ok := firstIdx[target - nums[i]]
		if ok {
			return []int{idx, i}
		} else {
			firstIdx[nums[i]] = i
		}
	}

	return []int{}
}
