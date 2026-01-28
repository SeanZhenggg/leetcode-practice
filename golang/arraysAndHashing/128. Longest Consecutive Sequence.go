package arraysAndHashing

import "sort"

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

// sort first + record max length for every consecutive subsequence
func longestConsecutive(nums []int) int {
	sort.Ints(nums)

	maxLength := 0
	length := 1
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
			if nums[i] == 1+nums[i-1] {
				length++
			} else {
				length = 1
			}
		}

		maxLength = max(maxLength, length)
	}
	return maxLength
}

// sort first + record max length for every consecutive subsequence
func longestConsecutive2(nums []int) int {
	sort.Ints(nums)

	maxLength := 0
	curr := 0
	length := 0
	for i := 0; i < len(nums); i++ {
		if curr != nums[i] {
			curr = nums[i]
			length = 0
		}

		// skip duplicate next value
		if i < len(nums)-1 && curr == nums[i+1] {
			continue
		}

		curr++
		length++

		maxLength = max(maxLength, length)
	}
	return maxLength
}

// hashmap + start from the leftmost
func longestConsecutive3(nums []int) int {
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
	}

	return maxLen
}

// hashmap + boundary
func longestConsecutive4(nums []int) int {
	// 用一個 hashmap 來記錄每一個數字所在的 sequence 的長度
	// 當前數字的 sequence 長度=左邊+右邊+1(自己)
	// 左邊代表的是 num 之前結束的 sequence 長度
	// 右邊代表的是 num 之後開始的 sequence 長度
	// m[i] = m[i-1]+m[i+1]+1
	// 更新左邊界跟右邊界的 sequence 長度
	// m[i-m[i-1]] = m[i]
	// m[i+m[i+1]] = m[i]
	// why? 因為已經算過的數字下次遇到就會跳過(有數字的代表已經有被算過), 而且我們算每一個數字在意的只有數字的左邊跟右邊
	// 更新最大的 sequence 長度到 maxLength
	m := make(map[int]int)

	maxLen := 0
	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}

		m[num] = m[num-1] + m[num+1] + 1 //更新當前數字的 sequence 長度
		m[num-m[num-1]] = m[num]         //更新左邊界的 sequence 長度
		m[num+m[num+1]] = m[num]         //更新右邊界的 sequence 長度

		maxLen = max(maxLen, m[num])
	}
	return maxLen
}
