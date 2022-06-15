/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var shiftGrid = function(grid, k) {
    let [m, n] = [grid.length, grid[0].length]
    let newGrid = Array.from({ length: m}).map(e => Array.from({ length: n}))

    for(var i = 0; i < m; i++) {
        for(var j = 0; j < n; j++) {
            let newY = (k + j) % n 
            let newX = (i + Math.floor((k + j) / n) ) % m
            
            newGrid[newX][newY] = grid[i][j]
        }
    }
    return newGrid
};