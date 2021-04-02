package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var seq []int
	seq = IntToSlice(x, seq)

	left := 0
	right := len(seq) - 1

	for left < right {
		if seq[left] == seq[right] {
			left = left + 1
			right = right - 1
		} else {
			break
		}
	}

	if left < right {
		return false
	}

	return true
}

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}
