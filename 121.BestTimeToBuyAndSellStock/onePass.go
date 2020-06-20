import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt32
	r := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if r < prices[i]-min {
			r = prices[i] - min
		}
	}
	return r
}
