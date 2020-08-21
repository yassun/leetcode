func isSubsequence(s string, t string) bool {
	sc, tc := 0, 0
	for sc != len(s) && tc != len(t) {
		if s[sc] == t[tc] {
			sc, tc = sc+1, tc+1
		} else {
			tc++
		}
	}

	return sc == len(s)
}
