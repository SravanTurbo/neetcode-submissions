func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
    
	var dfs func(int, []int, int)
	dfs = func(i int, current []int, remaining int) {
		if remaining == 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			res = append(res, temp)
			return
		}

		if remaining < 0 || i >= len(nums) {
			return
		}

		//include
		current = append(current, nums[i])
		dfs(i, current, remaining-nums[i])

		//exclude
		current = current[:len(current)-1]
		dfs(i+1, current, remaining)
	}

	dfs(0, []int{}, target)

	return res
}
