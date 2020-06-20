func maxProfit(prices []int) int {
	r := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			d := prices[j] - prices[i]
			if r < d {
				r = d
			}
		}
	}
	return r
}
