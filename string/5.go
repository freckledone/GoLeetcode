package string

func longestPalindrome(s string) string {
	result := ""
	maxLength := 0

	maxPalindrome := func(left int, right int) {
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
				result = s[left : right+1]
			}
			left -= 1
			right += 1
		}
	}

	for index := range s {
		maxPalindrome(index, index)
		maxPalindrome(index, index+1)
	}

	return result
}
