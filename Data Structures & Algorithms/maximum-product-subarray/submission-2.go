func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxi := math.MinInt32
	prevMax := 1
	prevMin := 1
	for _, n := range nums {
		curMax := max(n, max(prevMax * n, prevMin * n))
		curMin := min(n, min(prevMax * n, prevMin * n))

		maxi    = max(maxi, curMax)

		prevMax = curMax
		prevMin = curMin
	}

	return maxi
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
