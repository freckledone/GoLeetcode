package string

func longestCommonPrefix(strs []string) string {
	minStr, maxStr := minMax(strs)

	for index, ch := range minStr {
		if byte(ch) != maxStr[index] {
			return minStr[:index]
		}
	}

	return minStr
}

func minMax(strings []string) (string, string) {
	min := strings[0]
	max := strings[0]
	for _, str := range strings[1:] {
		if min >= str {
			min = str
		}

		if str >= max {
			max = str
		}
	}

	return min, max
}
