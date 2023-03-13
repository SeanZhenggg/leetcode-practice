class Solution {
    public boolean isValidSudoku(char[][] board) {
        Set<Character> showedBefore = new HashSet<>();
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[i][j] != '.' && !showedBefore.add(board[i][j])) return false;
            }
            showedBefore.clear();
        }
        showedBefore.clear();

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[j][i] != '.' && !showedBefore.add(board[j][i])) return false;
            }
            showedBefore.clear();
        }
        showedBefore.clear();

        for(int i = 0; i < board.length; i++) {
            for (int j = 0; j < 3; j++) {
                for (int k = 0; k < 3; k++) {
                    int x = (int) (3 * Math.floor(i / 3) + j);
                    int y = 3 * (i % 3) + k;
                    if(board[x][y] != '.' && !showedBefore.add(board[x][y])) return false;
                }
            }
            showedBefore.clear();
        }

        return true;
    }
}