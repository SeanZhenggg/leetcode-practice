package twoPointers

func maxArea(height []int) int {
	var l, r = 0, len(height) - 1
	var maxAr int
	var minH int
	for l <= r {
		if height[l] < height[r] {
			minH = height[l]
		} else {
			minH = height[r]
		}
		maxAr = max(maxAr, minH*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxAr
}
