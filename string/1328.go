package string

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	for index := range palindrome {
		if palindrome[index] != 'a' && !(len(palindrome)%2 == 1 && index == len(palindrome)/2) {

			return replaceAtIndex(palindrome, 'a', index)
		}
	}

	return replaceAtIndex(palindrome, 'b', len(palindrome)-1)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
