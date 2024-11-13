package medium

// wrong solution
func LongestConsecutiveWrong(nums []int) int {
	m := make(map[int]int)
	maxLength := 0
	for _, num := range nums {
		nextSmallerValue := num - 1
		nextBiggerValue := num + 1
		length := 1
		if _, found := m[nextSmallerValue]; found {
			length += m[nextSmallerValue]
			for m[nextSmallerValue] != 0 {
				m[nextSmallerValue] = length
				nextSmallerValue--
			}
		}
		if _, found := m[nextBiggerValue]; found {
			length += m[nextBiggerValue]
			for m[nextBiggerValue] != 0 {
				m[nextBiggerValue] = length
				nextBiggerValue++
			}
		}
		m[num] = length

		maxLength = max(maxLength, length)
	}

	return maxLength
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func LongestConsecutive(nums []int) int {
	m := make(map[int]bool)
	maxLen := 0
	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		length := 0
		if !m[num-1] {
			i := num
			for m[i] {
				length++
				i += 1
			}
			maxLen = max(maxLen, length)
		}

		//if !m[num+1] {
		//	for m[num-1] {
		//		length++
		//	}
		//	maxLen = max(maxLen, length)
		//}
	}

	return maxLen
}
