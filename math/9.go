package math

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {
	digits := digitCount(x)
	reversed := 0
	for 0 <= digits {
		reversed += (x % 10) * (int(math.Pow(10, float64(digits))))
		digits--
		x /= 10
	}

	return reversed
}

func digitCount(x int) int {
	count := 0
	for x != 0 {
		x /= 10
		count += 1
	}
	return count - 1
}
