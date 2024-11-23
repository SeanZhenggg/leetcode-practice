package easy

import "log"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if preI, found := m[num]; found {
			return []int{preI, i}
		}
		m[target-num] = i
	}
	return []int{} // unreachable
}

func Test_TwoSum() {
	ans1 := twoSum([]int{2, 7, 11, 15}, 9)
	log.Println("ans1: ", ans1)
	ans2 := twoSum([]int{3, 2, 4}, 6)
	log.Println("ans2: ", ans2)
	ans3 := twoSum([]int{3, 3}, 6)
	log.Println("ans3: ", ans3)
}
