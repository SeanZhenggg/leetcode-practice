package medium

import "log"

// l,r
// 0 1 2 3 4 5 6 7 8 9 10 11
// 3*4 = 12
// n
// 5 -> 1, 1
// 5 / 4, 5 % 4
// 1 * n + 1
// 1 * n + 3
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
