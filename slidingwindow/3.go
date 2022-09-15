package slidingwindow

func lengthOfLongestSubstring(s string) int {

	answer, i := 0, 0

	var mp = map[byte]int{}

	for j := 0; j < len(s); j++ {
		index, ok := mp[s[j]]

		if ok {
			answer = max(answer, j-i)
			for i <= index {
				delete(mp, s[i])
				i++
			}
			i = index + 1
		}
		mp[s[j]] = j
	}

	return max(answer, len(s)-i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
