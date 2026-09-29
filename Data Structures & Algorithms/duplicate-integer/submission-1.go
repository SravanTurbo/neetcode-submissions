func hasDuplicate(nums []int) bool {
    numsMap := map[int]bool{}

	for _, num := range nums {
		if numsMap[num] {
			return true
		}

		numsMap[num] = true
	}

	return false
}
