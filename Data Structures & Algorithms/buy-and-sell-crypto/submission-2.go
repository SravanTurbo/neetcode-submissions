/*
start: 11:18
*/
func maxProfit(prices []int) int {
	buyIdx, sellIdx := 0, 0
	maxPrft := 0

	for i:=1; i<len(prices); i++ {
		todayPrice := prices[i]

		if todayPrice > prices[sellIdx] {
			sellIdx = i
		}

		if todayPrice < prices[buyIdx] {
			buyIdx  = i
			sellIdx = i
		}

		maxPrft = max(maxPrft, prices[sellIdx]-prices[buyIdx])
	}

	return maxPrft
}
