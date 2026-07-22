func hasDuplicate(nums []int) bool {
    numMap := make(map[int]int)

	for i := range nums {
		numMap[nums[i]]+=1
		if numMap[nums[i]] == 2 {
			return true
		}
	}

	return false
}
