package easy

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, i := range nums {
		if _, found := m[i]; found {
			return true
		}
		m[i] = struct{}{}
	}

	return false
}
