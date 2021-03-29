import "math"

func myAtoi(str string) int {
	s := 0
	for s = 0; s < len(str); s++ {
		if !isDigitOrSign(str[s]) && str[s] != ' ' {
			return 0
		}

		if isDigitOrSign(str[s]) {
			break
		}
	}

	plus := true
	if s < len(str) {
		if str[s] == '-' {
			plus = false
			s++
		} else if str[s] == '+' {
			s++
			plus = true
		}
	}

	result := 0
	for ; s < len(str); s++ {
		if !isDigit(str[s]) {
			break
		}

		d := int(str[s]) - '0'
		var nResult int
		if !plus {
			d = -d
		}

		nResult = result*10 + d

		if nResult > math.MaxInt32 || nResult < math.MinInt32 || (nResult-d)/10 != result {
			if plus {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = nResult
	}

	return result
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

func isDigitOrSign(c byte) bool {
	if isDigit(c) || c == '-' || c == '+' {
		return true
	} else {
		return false
	}
}

