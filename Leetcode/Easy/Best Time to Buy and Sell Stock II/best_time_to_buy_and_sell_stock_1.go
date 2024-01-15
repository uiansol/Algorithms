// Accepted  5 ms  3.3 MB

func maxProfit(prices []int) int {
	lenght := len(prices)
	profit := 0

	for i := 0; i < lenght-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}

	return profit
}