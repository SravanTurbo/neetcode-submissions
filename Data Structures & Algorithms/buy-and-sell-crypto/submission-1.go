func maxProfit(prices []int) int {
	result := 0
	if len(prices)==1 {
		return result
	}

	l, r := 0, 1
	for r<len(prices){
		if prices[l]<prices[r]{
			result = max(result, prices[r]-prices[l])
		}else{
			l = r
		}
		r++
	}

	return result
}