package medium

func IsValidSudoku(board [][]byte) bool {
	var (
		rowNumCount = make(map[int]int)
		colNumCount = make(map[int]int)
		boxNumCount = make(map[int]int)
	)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			rowNum := getNumber(board, i, j)
			if rowNum != -1 {
				rowNumCount[rowNum]++
			}
			colNum := getNumber(board, j, i)
			if colNum != -1 {
				colNumCount[colNum]++
			}
			if rowNumCount[rowNum] > 1 || colNumCount[colNum] > 1 {
				return false
			}
		}
		rowNumCount = make(map[int]int)
		colNumCount = make(map[int]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				x, y := (i/3)*3+j, (i*3+k)%9
				boxNum := getNumber(board, x, y)
				if boxNum != -1 {
					boxNumCount[boxNum]++
				}
				if boxNumCount[boxNum] > 1 {
					return false
				}
			}
		}
		boxNumCount = make(map[int]int)
	}

	return true
}

func getNumber(board [][]byte, i, j int) int {
	if board[i][j] >= 48 && board[i][j] <= 57 {
		return int(board[i][j] - '0')
	} else {
		return -1
	}
}

func IsValidSudoku2(board [][]byte) bool {
	var (
		rowNumCount = make(map[int]int)
		colNumCount = make(map[int]int)
		boxNumCount = make(map[int]int)
	)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			rowNum := getNumber(board, i, j)
			if rowNum != -1 {
				rowNumCount[rowNum]++
			}
			colNum := getNumber(board, j, i)
			if colNum != -1 {
				colNumCount[colNum]++
			}
			if rowNumCount[rowNum] > 1 || colNumCount[colNum] > 1 {
				return false
			}
		}
		rowNumCount = make(map[int]int)
		colNumCount = make(map[int]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				x, y := (i/3)*3+j, 3*(i%3)+k // y is different from solution 1
				boxNum := getNumber(board, x, y)
				if boxNum != -1 {
					boxNumCount[boxNum]++
				}
				if boxNumCount[boxNum] > 1 {
					return false
				}
			}
		}
		boxNumCount = make(map[int]int)
	}

	return true
}
