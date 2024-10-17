package medium

type count struct {
	num   int
	count int
}

func TopKFrequent(nums []int, k int) []int {
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

	//log.Println("arr: ", arr)
	return result
}
