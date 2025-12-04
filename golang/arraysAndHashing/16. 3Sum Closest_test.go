package arraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	assertion := assert.New(t)

	nums1, target1 := []int{1, 3, -2, -5, 4, 8}, 6
	ans1 := threeSumClosest(nums1, target1)
	assertion.Equal(6, ans1)

	nums2, target2 := []int{-1, 2, 1, -4}, 1
	ans2 := threeSumClosest(nums2, target2)
	assertion.Equal(2, ans2)

	nums3, target3 := []int{0, 0, 0}, 1
	ans3 := threeSumClosest(nums3, target3)
	assertion.Equal(0, ans3)

	nums4, target4 := []int{1, 1, 1, 1}, -100
	ans4 := threeSumClosest(nums4, target4)
	assertion.Equal(3, ans4)
}
