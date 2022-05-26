package blind75

func maxProfit(prices []int) int {
	min := prices[0]
	max := 0
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			max = 0
		} else if prices[i] > max {
			max = prices[i]
			if profit < max-min {
				profit = max - min
			}
		}
	}
	return profit
}
