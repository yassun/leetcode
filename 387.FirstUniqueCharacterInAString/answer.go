func firstUniqChar(s string) int {
	m := map[string]int{}
	for _, v := range s {
		m[string(v)]++
	}
	for i, v := range s {
		if m[string(v)] == 1 {
			return i
		}
	}
	return -1
}
