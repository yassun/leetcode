/*
const MaxInt32 = int(2147483647)
const MinInt32 = int(-2147483648)
*/
import (
	"math"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		tmp := x % 10
		x = x / 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && tmp > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && tmp < -8) {
			return 0
		}
		result = result*10 + tmp
	}

	return result
}
