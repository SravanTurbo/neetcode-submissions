/*
idea: greedy (any number of transactions, buy and sell instead of hold)
*/
func maxProfit(prices []int) int {
	profit := 0 
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
