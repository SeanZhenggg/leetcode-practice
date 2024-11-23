package medium

import (
	"log"
	"slices"
)

type count struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))

	for _, n := range nums {
		m[n]++
	}

	arr := make([]count, 0, len(m))
	for num, c := range m {
		arr = append(arr, count{num, c})
	}

	// selection sort, can switch to other sorting algorithm like quick sort(n * logn)
	var maxIndex = 0
	for i := 0; i < len(arr); i++ {
		maxIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[maxIndex].count < arr[j].count {
				maxIndex = j
			}
		}

		arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
	}

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, arr[i].num)
	}

	return result
}

func topKFrequentReview1(nums []int, k int) []int {
	m := make(map[int]int, len(nums))

	for _, v := range nums {
		m[v]++
	}

	counts := make([]count, 0, len(m))
	for num, c := range m {
		counts = append(counts, count{num, c})
	}

	slices.SortFunc(counts, func(i, j count) int {
		return j.count - i.count
	})

	log.Printf("counts: %v", counts)

	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, counts[i].num)
	}

	return ret
}

func Test_TopKFrequent() {
	ans1 := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Println("ans1: ", ans1)
	ans2 := topKFrequent([]int{1}, 1)
	log.Println("ans2: ", ans2)
}

func Test_TopKFrequentReview1() {
	ans1 := topKFrequentReview1([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Println("ans1: ", ans1)
	ans2 := topKFrequentReview1([]int{1}, 1)
	log.Println("ans2: ", ans2)
	ans3 := topKFrequentReview1([]int{-5, -2, -1, 2, 5, -5, 2, -1}, 3)
	log.Println("ans3: ", ans3)
}
