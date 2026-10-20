/*
start: 11:18
*/
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}

	l, r := 0, 1
	maxP := 0

	for r < n {
		if prices[r] > prices[l] {
			maxP = max(maxP, prices[r]-prices[l])
		} else {
			l = r
		}
		r++
	}

	return maxP
}
