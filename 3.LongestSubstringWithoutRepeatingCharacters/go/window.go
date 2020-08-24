func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]bool)
	max, right, left := 0, 0, 0
	for left < len(s) && right < len(s) {
		// windowに存在する場合は無くなるまでleftを移動
		if _, okay := window[s[right]]; okay == true {
			delete(window, s[left])
			left = left + 1
			// windowに存在しない場合はwindowに格納してrightを移動
		} else {
			window[s[right]] = true
			right = right + 1
			if max < len(window) {
				max = len(window)
			}
		}
	}
	return max
}
