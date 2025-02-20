package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"log"
	"math/rand"
	"sort"
	"time"
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

	for l <= r {
		pivot := partition(nums, l, r)
		if pivot == kthLargest {
			return nums[pivot]
		} else if pivot > kthLargest {
			r = pivot - 1
			//pivot = partition(nums, l, pivot-1) // wrong
		} else {
			l = pivot + 1
			//pivot = partition(nums, pivot+1, r) // wrong
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
			nums[j+1], nums[i] = nums[i], nums[j+1]
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

// 5. other's solution
func findKthLargest5(nums []int, k int) int {
	rand.New(rand.NewSource(time.Now().UnixNano())) // https://gabrieleromanato.name/go-rand-seed-has-been-deprecated-a-solution
	index := quickSelect(nums, 0, len(nums)-1, k)
	return nums[index]
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return left
	}

	pivotIndex := rand.Intn(right-left+1) + left
	pivot := nums[pivotIndex]

	// 三路切分: lt 表示小於 pivot 的右邊界
	//          gt 表示大於 pivot 的左邊界
	lt, i, gt := left, left, right
	for i <= gt {
		if nums[i] > pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] < pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	if k-1 < lt {
		return quickSelect(nums, left, lt-1, k)
	}
	if k-1 > gt {
		return quickSelect(nums, gt+1, right, k)
	}
	return lt
}

// 6. other's solution
func findKthLargest6(nums []int, k int) int {
	n := len(nums)
	r := rand.Intn(n)

	nums[r], nums[n-1] = nums[n-1], nums[r]

	pivot := nums[n-1]

	l, r, u := 0, 0, n-2
	for r <= u {
		switch {
		case nums[r] > pivot:
			nums[l], nums[r] = nums[r], nums[l]
			r++
			l++
		case nums[r] < pivot:
			nums[r], nums[u] = nums[u], nums[r]
			u--
		default:
			r++
		}
	}

	nums[r], nums[n-1] = nums[n-1], nums[r]

	if r < k-1 {
		return findKthLargest(nums[r+1:], k-r-1)
	}

	if l > k-1 {
		return findKthLargest(nums[:l], k)
	}

	return nums[k-1]
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
	nums1 := []int{1, 2, 3, 4, 5, 6}
	k1 := 1
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
