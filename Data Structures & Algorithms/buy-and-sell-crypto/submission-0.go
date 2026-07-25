func maxProfit(prices []int) int {
	bestBuy, profit := 101, 0
	for i:=1; i<len(prices); i++{ //selling
		bestBuy = min(bestBuy, prices[i-1])
		if prices[i]>bestBuy{
			profit = max(profit, prices[i]-bestBuy)
		}
	}

	return profit
}
