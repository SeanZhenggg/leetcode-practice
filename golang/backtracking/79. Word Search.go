package backtracking

func exist(board [][]byte, word string) bool {
	visited := make(map[[2]int]bool)

	var dfs func(x, y int, index int) bool
	dfs = func(x, y int, index int) bool {
		visited[[2]int{x, y}] = true
		defer func() {
			visited[[2]int{x, y}] = false
		}()
		// not found/found
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}

		// up
		if x-1 >= 0 && !visited[[2]int{x - 1, y}] {
			if dfs(x-1, y, index+1) {
				return true
			}
		}
		// down
		if x+1 < len(board) && !visited[[2]int{x + 1, y}] {
			if dfs(x+1, y, index+1) {
				return true
			}
		}
		// left
		if y-1 >= 0 && !visited[[2]int{x, y - 1}] {
			if dfs(x, y-1, index+1) {
				return true
			}
		}
		// right
		if y+1 < len(board[x]) && !visited[[2]int{x, y + 1}] {
			if dfs(x, y+1, index+1) {
				return true
			}
		}

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
