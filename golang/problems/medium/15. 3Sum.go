package medium

import (
	"log"
	"slices"
)

// first solution, sort + two pointers, O(n^2)
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0, len(nums))
	dup := make(map[[3]int]struct{})

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		left, right := 0, len(nums)-1

		diff := 0 - nums[i]

		for left < right {
			if left == i {
				left++
				continue
			}
			if right == i {
				right--
				continue
			}
			sum := nums[left] + nums[right]
			if sum < diff {
				left++
			} else if sum > diff {
				right--
			} else {
				k := genKey(nums[left], nums[right], nums[i])
				if _, found := dup[k]; !found {
					ret = append(ret, []int{nums[left], nums[right], nums[i]})
					dup[k] = struct{}{}
				}
				left++
				right--
			}
		}
	}

	return ret
}

func genKey(x, y, z int) [3]int {
	var maxV int = x
	var minV int = x
	var midV int

	if y > maxV {
		maxV = y
	}
	if z > maxV {
		maxV = z
	}

	if y < minV {
		minV = y
	}
	if z < minV {
		minV = z
	}

	if maxV == x && minV == y || (maxV == y && minV == x) {
		midV = z
	}
	if maxV == y && minV == z || (maxV == z && minV == y) {
		midV = x
	}
	if maxV == x && minV == z || (maxV == z && minV == x) {
		midV = y
	}

	return [3]int{minV, midV, maxV}
}

// second solution, sort + two pointers, better one, O(n^2)
func threeSum2(nums []int) [][]int {
	ret := make([][]int, 0, len(nums))

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ret = append(ret, []int{nums[l], nums[r], nums[i]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return ret
}

func threeSumReview(nums []int) [][]int {
	ret := make([][]int, 0, len(nums))
	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		//duplicated, calculated before, just ignore
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]

			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ret = append(ret, []int{nums[l], nums[r], nums[i]})

				l++
				// ignore duplicated l+1
				for l < r && nums[l-1] == nums[l] {
					l++
				}
			}
		}
	}

	return ret
}

func Test_ThreeSum() {
	case1 := []int{-1, 0, 1, 2, -1, -4}
	ans1 := threeSum(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{0, 1, 1}
	ans2 := threeSum(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{0, 0, 0}
	ans3 := threeSum(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	ans4 := threeSum(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_ThreeSum2() {
	case1 := []int{-1, 0, 1, 2, -1, -4}
	ans1 := threeSum2(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{0, 1, 1}
	ans2 := threeSum2(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{0, 0, 0}
	ans3 := threeSum2(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	ans4 := threeSum2(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_ThreeSumReview() {
	case1 := []int{-1, 0, 1, 2, -1, -4}
	ans1 := threeSumReview(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{0, 1, 1}
	ans2 := threeSumReview(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{0, 0, 0}
	ans3 := threeSumReview(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	ans4 := threeSumReview(case4)
	log.Printf("ans4: %v", ans4)
}
