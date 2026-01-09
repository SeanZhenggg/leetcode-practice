package arraysAndHashing

// solution 1
func productExceptSelf(nums []int) []int {
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
func productExceptSelf2(nums []int) []int {
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
func productExceptSelf3(nums []int) []int {
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

func productExceptSelf4(nums []int) []int {
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			leftProduct[i] = 1
		} else {
			leftProduct[i] = leftProduct[i-1] * nums[i-1]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightProduct[i] = 1
		} else {
			rightProduct[i] = rightProduct[i+1] * nums[i+1]
		}
	}

	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = leftProduct[i] * rightProduct[i]
	}

	return ret
}

func productExceptSelf5(nums []int) []int {
	prevProduct := 1
	ret := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ret[i] = prevProduct
		prevProduct *= nums[i]
	}

	prevProduct = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= prevProduct
		prevProduct *= nums[i]
	}

	return ret
}

// review
func productExceptSelfReview(nums []int) []int {
	var leftProducts = make([]int, len(nums))
	var rightProducts = make([]int, len(nums))
	var lProduct = 1
	var rProduct = 1
	for i := 0; i < len(nums); i++ {
		leftProducts[i] = lProduct
		lProduct *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rightProducts[i] = rProduct
		rProduct *= nums[i]
	}

	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = leftProducts[i] * rightProducts[i]
	}

	return ret
}

func productExceptSelfReview2(nums []int) []int {
	var ret = make([]int, len(nums))

	var prod = 1
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
