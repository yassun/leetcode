package main

func romanToInt(s string) int {
	romanValue := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	relust := 0
	i := 0
	for i < len(s) {
		if i+1 < len(s) {
			if romanValue[s[i]] < romanValue[s[i+1]] {
				relust += romanValue[s[i+1]] - romanValue[s[i]]
				i += 2
				continue
			} else {
				relust += romanValue[s[i]]
			}
		} else {
			relust += romanValue[s[i]]

		}
		i += 1
	}
	return relust
}
