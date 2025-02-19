package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"log"
	"math/rand"
	"sort"
)

// 1. min heap solution
func findKthLargest(nums []int, k int) int {
	pq := priorityqueue.NewWith(utils.IntComparator)

	for i := 0; i < len(nums); i++ {
		pq.Enqueue(nums[i])

		if pq.Size() > k {
			pq.Dequeue()
		}
	}
	val, _ := pq.Peek()
	return val.(int)
}

// 2. quick select solution(doesn't work anymore, caused test case has many duplicated values)
func findKthLargest2(nums []int, k int) int {
	kthLargest := len(nums) - k
	l, r := 0, len(nums)-1
	pivot := partition(nums, l, r)

	for l <= r {
		if pivot == kthLargest {
			return nums[pivot]
		} else if pivot > kthLargest {
			pivot = partition(nums, l, pivot-1)
		} else {
			pivot = partition(nums, pivot+1, r)
		}
	}

	return 0 // unreachable
}

func partition(nums []int, l, r int) int {
	pivot := int(float64(r-l+1)*rand.Float64() + float64(l))
	// IMPORTANT: swap the random pivot index for end index to make the algo below still working
	nums[r], nums[pivot] = nums[pivot], nums[r]
	pivot = r
	j := l - 1
	for i := l; i < r; i++ {
		if nums[i] <= nums[pivot] {
			j++
		}
	}

	nums[j+1], nums[pivot] = nums[pivot], nums[j+1]
	return j + 1
}

// 3. quick select solution(hoare's partitioning)
func findKthLargest3(nums []int, k int) int {
	l, r := 0, len(nums)-1
	var pivot int
	for l <= r {
		pivot = partition3(nums, l, r)
		if pivot > k-1 {
			r = pivot - 1
		} else if pivot < k-1 {
			l = pivot + 1
		} else {
			return nums[pivot]
		}
	}
	return nums[pivot]
}

func partition3(nums []int, l, r int) int {
	pivot := nums[r]
	i, j := l, r-1

	for i <= j {
		for i < r && nums[i] > pivot {
			i++
		}
		for j > l && nums[j] < pivot {
			j--
		}
		// i <= pivot, pivot <= j
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			break
		}
	}

	nums[i], nums[r] = nums[r], nums[i]
	return i
}

// 4. sort solution
func findKthLargest4(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func Test_FindKthLargest() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	ans1 := findKthLargest(nums1, k1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	ans2 := findKthLargest(nums2, k2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{3, 2, 1, 5, 6, 4}
	k3 := 4
	ans3 := findKthLargest(nums3, k3)
	log.Printf("ans3: %v", ans3)

}

func Test_FindKthLargest2() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	ans1 := findKthLargest2(nums1, k1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	ans2 := findKthLargest2(nums2, k2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{3, 2, 1, 5, 6, 4}
	k3 := 4
	ans3 := findKthLargest2(nums3, k3)
	log.Printf("ans3: %v", ans3)

}

func Test_FindKthLargest3() {
	nums1 := []int{3, 6, 5, 4, 2, 1}
	k1 := 2
	ans1 := findKthLargest3(nums1, k1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	ans2 := findKthLargest3(nums2, k2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{3, 2, 1, 5, 6, 4}
	k3 := 4
	ans3 := findKthLargest3(nums3, k3)
	log.Printf("ans3: %v", ans3)

}
