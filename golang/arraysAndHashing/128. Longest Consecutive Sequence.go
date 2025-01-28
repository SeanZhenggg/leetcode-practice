package arraysAndHashing

import "log"

// wrong solution
func longestConsecutiveWrong(nums []int) int {
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

func longestConsecutive(nums []int) int {
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

func longestConsecutiveReview(nums []int) int {
	m := make(map[int]bool)
	maxLength := 0
	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		if !m[num-1] {
			i := 1
			for m[num+1] {
				i++
				num += 1
			}
			if i > maxLength {
				maxLength = i
			}
		}
	}

	return maxLength
}

func Test_LongestConsecutive() {
	ans1 := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	log.Println("ans1: ", ans1)
	ans2 := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	log.Println("ans2: ", ans2)
	ans3 := longestConsecutive([]int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161})
	log.Println("ans3: ", ans3)
}

func Test_LongestConsecutiveReview() {
	ans1 := longestConsecutiveReview([]int{100, 4, 200, 1, 3, 2})
	log.Println("ans1: ", ans1)
	ans2 := longestConsecutiveReview([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	log.Println("ans2: ", ans2)
	ans3 := longestConsecutiveReview([]int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161})
	log.Println("ans3: ", ans3)
}
