package binarySearch

import "log"

// binary solution - O(log(m*n))
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, (m*n)-1

	for l <= r {
		mid := l + (r-l)/2
		x, y := mid/n, mid%n
		if matrix[x][y] > target {
			r = mid - 1
		} else if matrix[x][y] == target {
			return true
		} else {
			l = mid + 1
		}
	}
	return false
}

// binary solution - O(log(m)+log(n))
func searchMatrix2(matrix [][]int, target int) bool {
	t, b := 0, len(matrix)-1
	for t <= b {
		mid := t + (b-t)/2
		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
			break
		} else if matrix[mid][0] > target {
			b = mid - 1
		} else if matrix[mid][len(matrix[mid])-1] < target {
			t = mid + 1
		}
	}
	if t > len(matrix)-1 || b < 0 {
		return false
	}
	row := t + (b-t)/2
	l, r := 0, len(matrix[row])-1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[row][mid] < target {
			l = mid + 1
		} else if matrix[row][mid] == target {
			return true
		} else {
			r = mid - 1
		}
	}
	return false
}

func Test_searchMatrix() {
	matrix1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	t1 := 3
	ans1 := searchMatrix(matrix1, t1)
	log.Printf("ans1: %v", ans1)
	matrix2 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	t2 := 13
	ans2 := searchMatrix(matrix2, t2)
	log.Printf("ans2: %v", ans2)
}

func Test_searchMatrix2() {
	//matrix1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	//t1 := 3
	//ans1 := searchMatrix2(matrix1, t1)
	//log.Printf("ans1: %v", ans1)
	//matrix2 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	//t2 := 13
	//ans2 := searchMatrix2(matrix2, t2)
	//log.Printf("ans2: %v", ans2)
	matrix3 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	t3 := 0
	ans3 := searchMatrix2(matrix3, t3)
	log.Printf("ans3: %v", ans3)
}
