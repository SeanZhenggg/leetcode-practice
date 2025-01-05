package hard

import "log"

// merge-sort solution, O(n+m)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := merge(nums1, nums2)
	mid := len(mergedArr) / 2
	if len(mergedArr)%2 != 0 {
		return float64(mergedArr[mid])
	} else {
		return (float64(mergedArr[mid-1]) + float64(mergedArr[mid])) / float64(2)
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if (n1+n2)%2 == 0 {
		a := float64(recurBinarySearch(nums1, 0, n1-1, nums2, 0, n2-1, (n1+n2)/2))
		b := float64(recurBinarySearch(nums1, 0, n1-1, nums2, 0, n2-1, (n1+n2)/2-1))
		return (a + b) / 2
	}

	return float64(recurBinarySearch(nums1, 0, n1-1, nums2, 0, n2-1, (n1+n2)/2))
}

func recurBinarySearch(nums1 []int, aStart, aEnd int, nums2 []int, bStart, bEnd int, k int) int {
	if aStart > aEnd {
		return nums2[k-aStart]
	}
	if bStart > bEnd {
		return nums1[k-bStart]
	}

	aMid, bMid := (aStart+aEnd)/2, (bStart+bEnd)/2
	aMidVal, bMidVal := nums1[aMid], nums2[bMid]

	if aMid+bMid < k {
		// kth element is in the right half of total elements
		if aMidVal < bMidVal {
			return recurBinarySearch(nums1, aMid+1, aEnd, nums2, bStart, bEnd, k)
		} else {
			return recurBinarySearch(nums1, aStart, aEnd, nums2, bMid+1, bEnd, k)
		}
	} else {
		// kth element is in the left half of total elements
		if aMidVal < bMidVal {
			return recurBinarySearch(nums1, aStart, aEnd, nums2, bStart, bMid-1, k)
		} else {
			return recurBinarySearch(nums1, aStart, aMid-1, nums2, bStart, bEnd, k)
		}
	}
}

func merge(nums1 []int, nums2 []int) []int {
	l, r := 0, 0
	ret := make([]int, 0, len(nums1)+len(nums2))
	for l < len(nums1) && r < len(nums2) {
		if nums1[l] < nums2[r] {
			ret = append(ret, nums1[l])
			l++
		} else {
			ret = append(ret, nums2[r])
			r++
		}
	}

	for l < len(nums1) {
		ret = append(ret, nums1[l])
		l++
	}

	for r < len(nums2) {
		ret = append(ret, nums2[r])
		r++
	}

	return ret
}

func Test_findMedianSortedArrays() {
	nums11 := []int{1, 3}
	nums12 := []int{2}
	ans1 := findMedianSortedArrays(nums11, nums12)
	log.Printf("ans1: %v", ans1)
	nums21 := []int{1, 2}
	nums22 := []int{3, 4}
	ans2 := findMedianSortedArrays(nums21, nums22)
	log.Printf("ans2: %v", ans2)

}

func Test_findMedianSortedArrays2() {
	nums11 := []int{1, 3}
	nums12 := []int{2}
	ans1 := findMedianSortedArrays2(nums11, nums12)
	log.Printf("ans1: %v", ans1)
	nums21 := []int{1, 2}
	nums22 := []int{3, 4}
	ans2 := findMedianSortedArrays2(nums21, nums22)
	log.Printf("ans2: %v", ans2)

}
