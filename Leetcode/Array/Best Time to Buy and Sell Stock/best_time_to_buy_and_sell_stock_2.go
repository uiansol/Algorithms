// Accepted  99 ms  8.2 MB

func maxProfit(prices []int) int {
	min_price, max_profit := 1000000, 0

	for i := range prices {
		if prices[i] < min_price {
			min_price = prices[i]
		} else if prices[i]-min_price > max_profit {
			max_profit = prices[i] - min_price
		}
	}

	return max_profit
}