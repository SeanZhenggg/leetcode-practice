package arraysAndHashing

import "github.com/emirpasic/gods/sets/hashset"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < i+1+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			if i-v <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

func containsNearbyDuplicate3(nums []int, k int) bool {
	s := hashset.New()
	for i := 0; i < len(nums); i++ {
		if s.Contains(nums[i]) {
			return true
		}
		s.Add(nums[i])
		if s.Size() > k {
			s.Remove(nums[i-k])
		}
	}
	return false
}
