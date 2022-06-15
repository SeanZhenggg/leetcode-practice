/**
 * @param {number[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var gameOfLife = function(board) {
    let dieOrLive = {}
    let [m, n] = [board.length, board[0].length]
    for(let i = 0; i < m ; i ++) {
        for(let j = 0; j < n; j ++) {

            let liveCount = 0;
            console.log('(i,j)', i,j)
            for(let k = i-1; k<=i+1; k++) {
                for(let l = j-1; l<=j+1; l++) {
                    if(k < 0 || k > m-1) continue
                    if(l < 0 || l > n-1) continue
                    
                    if(board[k][l] === 1) {
                        if(k !== i || l !== j) liveCount ++;
                    }
                }   
            }
            if(!board[i][j]) {
                dieOrLive[i*n+j] = liveCount === 3 ? 1 : 0
            } 
            else {
                if(liveCount > 3) {
                    dieOrLive[i*n+j] = 0
                }
                else if(2 <= liveCount && liveCount <= 3) {
                    dieOrLive[i*n+j] = 1
                } 
                else {
                    dieOrLive[i*n+j] = 0
                }
            }
        }
    }

    console.log('dieOrLive', dieOrLive)
    
    for(let o = 0; o < m * n; o++) {
        board[Math.floor(o / n)][o % n]= dieOrLive[o]
    }
};