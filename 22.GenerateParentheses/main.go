package main

func rec(p string, left int, right int, max int, result *[]string) {
	if left < right {
		return
	}
	if left == max && right == max {
		if len(p) > 0 {
			*result = append(*result, p)
		}
	} else {
		if left < max {
			rec(p+"(", left+1, right, max, result)
		}
		if right < left {
			rec(p+")", left, right+1, max, result)
		}
	}
}
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	rec("", 0, 0, n, &result)
	return result
}
