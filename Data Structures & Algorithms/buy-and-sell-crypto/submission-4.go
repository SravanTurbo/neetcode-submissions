/*
start: 11:50
idea: greedy
t: O(n), s: O(1)
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice  := prices[0]
	maxProfit := 0

	for i:=1; i<len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		maxProfit = max(maxProfit, prices[i] - minPrice)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

