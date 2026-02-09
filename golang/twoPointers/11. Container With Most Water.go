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

func maxAreaReview(height []int) int {
	l, r := 0, len(height)-1
	maxVolume := 0
	// we don't need to concern about l == r, because it will cause the width to be 0
	for l < r {
		minHeight := height[l]
		if minHeight > height[r] {
			minHeight = height[r]
		}

		volume := (r - l) * minHeight

		if maxVolume < volume {
			maxVolume = volume
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxVolume
}
