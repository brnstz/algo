package algo

// MaxProfit accepts a series of prices over time and returns the maximum
// profit that can be extracted by buying at time i and selling at time j where
// j > i.
func MaxProfit(prices []int) int {
	var (
		minBuy    = MaxIntVal
		maxProfit = MinIntVal
		profit    int
	)

	for _, price := range prices {
		profit = price - minBuy

		if profit > maxProfit {
			maxProfit = profit
		}

		if price < minBuy {
			minBuy = price
		}
	}

	return maxProfit
}
