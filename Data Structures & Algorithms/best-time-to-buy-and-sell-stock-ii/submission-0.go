func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	l, r := 0, 0 //buy, sell

	for r < len(prices)-1 {
		if prices[r+1] >=  prices[r] {
			r++ //hold
		} else {
			if prices[r] - prices[l] > 0 { //sell
				profit += prices[r] - prices[l]
			}
	
			r++
			l=r //next-buy
		}
	}

	if prices[r] - prices[l] > 0 { //final-sell
		profit += prices[r] - prices[l]
	}

	return profit
}
