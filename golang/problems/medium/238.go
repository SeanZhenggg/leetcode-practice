package medium

// solution 1
func ProductExceptSelf(nums []int) []int {
	leftProducts := make([]int, len(nums))
	rightProducts := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		leftOne := 1
		if i-1 >= 0 {
			leftOne = leftProducts[i-1]
		}
		leftProducts[i] = leftOne * nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		rightOne := 1
		if i+1 <= len(nums)-1 {
			rightOne = rightProducts[i+1]
		}
		rightProducts[i] = rightOne * nums[i]
	}

	ret := make([]int, 0, len(nums))

	for i := range nums {
		left := 1
		right := 1

		if i-1 >= 0 {
			left *= leftProducts[i-1]
		}

		if i+1 <= len(nums)-1 {
			right *= rightProducts[i+1]
		}

		ret = append(ret, left*right)
	}

	return ret
}

// solution 2 - using O(1)
func ProductExceptSelf2(nums []int) []int {
	ret := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		left := 1
		leftNum := 1
		if i-1 >= 0 {
			left = ret[i-1]
			leftNum = nums[i-1]
		}
		ret[i] = left * leftNum
	}

	currentRightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= currentRightProduct
		currentRightProduct *= nums[i]
	}

	return ret
}

// solution - using O(1) with more elegant code
func ProductExceptSelf3(nums []int) []int {
	ret := make([]int, len(nums))
	prod := 1

	for i := 0; i < len(nums); i++ {
		ret[i] = prod
		prod *= nums[i]
	}

	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= prod
		prod *= nums[i]
	}

	return ret
}
