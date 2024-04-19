package besttimebysell

func MaxProfit(prices []int) int {
	profit := 0
	byPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < byPrice {
			byPrice = prices[i]
		} else if prices[i]-byPrice > profit {
			profit = prices[i] - byPrice
		}
	}
	return profit
}

func MaxProfitLevel2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
