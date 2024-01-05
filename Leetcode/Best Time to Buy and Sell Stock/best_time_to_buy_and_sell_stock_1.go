// Time Limit Exceeded

func maxProfit(prices []int) int {
	l, maxProfit, profit := len(prices), 0, 0

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			profit = prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}