package easy

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if preI, found := m[num]; found {
			return []int{preI, i}
		}
		m[target-num] = i
	}
	return []int{} // unreachable
}
