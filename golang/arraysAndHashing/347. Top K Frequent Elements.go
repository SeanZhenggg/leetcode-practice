package arraysAndHashing

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"math/rand"
	"slices"
)

type count struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
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

	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, counts[i].num)
	}

	return ret
}

func topKFrequent2(nums []int, k int) []int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		aval, bval := a.(count), b.(count)
		if aval.count < bval.count {
			return 1
		} else if aval.count > bval.count {
			return -1
		} else {
			return 0
		}
	})

	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		pq.Enqueue(count{k, v})
	}

	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		val, ok := pq.Dequeue()
		if ok {
			ret = append(ret, val.(count).num)
		}
	}
	return ret
}

func topKFrequent3(nums []int, k int) []int {
	countMap := make(map[int]int, len(nums))
	for _, i2 := range nums {
		countMap[i2]++
	}

	unique := make([]int, 0, len(countMap))
	for k := range countMap {
		unique = append(unique, k)
	}
	n := len(unique)

	_quickSelect(countMap, unique, 0, n-1, n-k)

	ret := make([]int, k)
	copy(ret, unique[n-k:n])
	return ret
}

func _quickSelect(countMap map[int]int, unique []int, l, r int, k int) {
	pivot := l + rand.Intn(r-l+1)

	idx := _partition(countMap, unique, l, r, pivot)
	if idx > k {
		_quickSelect(countMap, unique, l, idx-1, k)
	} else if idx < k {
		_quickSelect(countMap, unique, idx+1, r, k)
	} else {
		return
	}
}

// 1 4 2 3 5
// 1 4 5 3 2
// 1
// 1 4
// 1 4 5 3 2
// 1 2 5 3 4
// 5 3 4
// 5 4 3
// 3 4 5
func _partition(countMap map[int]int, unique []int, l, r int, pivot int) int {
	freq := countMap[unique[pivot]]

	unique[pivot], unique[r] = unique[r], unique[pivot]
	idx := l
	for i := l; i < r; i++ {
		if countMap[unique[i]] < freq {
			unique[idx], unique[i] = unique[i], unique[idx]
			idx++
		}
	}
	unique[idx], unique[r] = unique[r], unique[idx]
	return idx
}

func topKFrequent4(nums []int, k int) []int {
	// 記錄每個數字出現的次數
	// 建立一個雙層陣列, 陣列的索引用來表示出現次數, 對應的資料為所有出現次數相同的數字
	// 這樣子之後從尾巴遍歷, 就意味著是從出現次數高到底的順序依序遍歷
	// 從陣列的尾巴開始遍歷, 當索引內有值則取出值並記錄目前取出的數量
	// 當取出的數量 = k 時, 代表已經取完前 k 個出現次數最多的數字

	countMap := make(map[int]int, len(nums))
	for _, num := range nums {
		countMap[num]++
	}

	valuesOfCount := make([][]int, len(nums)+1)
	for num, count := range countMap {
		valuesOfCount[count] = append(valuesOfCount[count], num)
	}

	ret := make([]int, 0, k)
LOOP:
	for i := len(valuesOfCount) - 1; i >= 0; i-- {
		for _, num := range valuesOfCount[i] {
			ret = append(ret, num)
			if len(ret) == k {
				break LOOP
			}
		}
	}

	return ret
}
