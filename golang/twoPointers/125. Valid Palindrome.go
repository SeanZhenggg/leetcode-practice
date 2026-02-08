package twoPointers

// solution 1
func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	var (
		left  = 0
		right = len(s) - 1
	)

	for left <= right {
		if isNotAlphaNumeric(s[left]) {
			left++
			continue
		}
		if isNotAlphaNumeric(s[right]) {
			right--
			continue
		}
		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isNotAlphaNumeric(c uint8) bool {
	return !(c >= 48 && c <= 57) && !(c >= 65 && c <= 90) && !(c >= 97 && c <= 122)
}

func toLowerCase(c uint8) uint8 {
	if c < 65 || c > 90 {
		return c
	}

	return c + 32
}
