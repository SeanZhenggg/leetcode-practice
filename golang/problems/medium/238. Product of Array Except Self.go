package medium

import "log"

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

func Test_ProductExceptSelf() {
	ans1 := productExceptSelf([]int{1, 2, 3, 4})
	log.Println("ans1: ", ans1)
	ans2 := productExceptSelf([]int{-1, 1, 0, -3, 3})
	log.Println("ans2: ", ans2)
}

func Test_ProductExceptSelf2() {
	ans3 := productExceptSelf2([]int{1, 2, 3, 4})
	log.Println("ans3: ", ans3)
	ans4 := productExceptSelf2([]int{-1, 1, 0, -3, 3})
	log.Println("ans4: ", ans4)
}

func Test_ProductExceptSelf3() {
	ans5 := productExceptSelf3([]int{1, 2, 3, 4})
	log.Println("ans5: ", ans5)
	ans6 := productExceptSelf3([]int{-1, 1, 0, -3, 3})
	log.Println("ans6: ", ans6)
}

func Test_ProductExceptSelfReview() {
	ans1 := productExceptSelfReview([]int{1, 2, 3, 4})
	log.Println("ans1: ", ans1)
	ans2 := productExceptSelfReview([]int{-1, 1, 0, -3, 3})
	log.Println("ans2: ", ans2)
	ans3 := productExceptSelfReview([]int{8, 2, -1, 3, 5})
	log.Println("ans3: ", ans3)
}

func Test_ProductExceptSelfReview2() {
	ans1 := productExceptSelfReview2([]int{1, 2, 3, 4})
	log.Println("ans1: ", ans1)
	ans2 := productExceptSelfReview2([]int{-1, 1, 0, -3, 3})
	log.Println("ans2: ", ans2)
	ans3 := productExceptSelfReview2([]int{8, 2, -1, 3, 5})
	log.Println("ans3: ", ans3)
}
