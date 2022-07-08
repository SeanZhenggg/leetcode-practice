/**
 * @param {character[][]} board
 * @return {boolean}
 */
var isValidSudoku = function (board) {
    let set = new Set()
    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            if (board[i][j] === '.') continue
            if (set.has(board[i][j])) return false
            else set.add(board[i][j])
        }
        set.clear()
    }
    set.clear()

    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board.length; j++) {
            if (board[j][i] === '.') continue
            if (set.has(board[j][i])) return false
            else set.add(board[j][i])
        }
        set.clear()
    }
    set.clear()

    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < 3; j++) {
            for (let k = 0; k < 3; k++) {
                let x = Math.floor(i / 3) * 3 + j, y = (i % 3) * 3 + k
                if (board[x][y] === '.') continue
                if (set.has(board[x][y])) return false
                else set.add(board[x][y])
            }
        }
        set.clear()
    }

    return true;
};