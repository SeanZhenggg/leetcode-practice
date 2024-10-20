package easy

// solution 1
func IsPalindrome(s string) bool {
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
		if transToLowercase(s[left]) != transToLowercase(s[right]) {
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

func transToLowercase(c uint8) uint8 {
	if c < 65 || c > 90 {
		return c
	}

	return c + 32
}

// solution 2 - without functions
func IsPalindrome2(s string) bool {
	if len(s) == 1 {
		return true
	}

	var (
		left  = 0
		right = len(s) - 1
	)

	for left <= right {
		if !(s[left] >= 48 && s[left] <= 57) && !(s[left] >= 65 && s[left] <= 90) && !(s[left] >= 97 && s[left] <= 122) {
			left++
			continue
		}
		if !(s[right] >= 48 && s[right] <= 57) && !(s[right] >= 65 && s[right] <= 90) && !(s[right] >= 97 && s[right] <= 122) {
			right--
			continue
		}
		l := s[left]
		r := s[right]

		if l >= 65 && l <= 90 {
			l += 32
		}
		if r >= 65 && r <= 90 {
			r += 32
		}
		if l != r {
			return false
		}

		left++
		right--
	}

	return true
}
