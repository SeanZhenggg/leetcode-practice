/**
 * @param {number} n
 * @return {number[][]}
 */
var generateMatrix = function (n) {
  let output = Array.from({ length: n }).map(e => Array.from({ length: n }))
  let x = 0, y = 0
  let direction = 'r'

  let moveByDir = {
    r: () => x++,
    d: () => y++,
    l: () => x--,
    u: () => y--
  }

  let yAxisInsideBorder = y => y >= 0 && y <= n - 1
  let xAxisInsideBorder = x => x >= 0 && x <= n - 1
  let isAlreadyWentThrough = (x, y) => output[y][x] !== undefined

  output[y][x] = 1
  moveByDir[direction]()

  for (let i = 2; i <= n * n; i++) {
    switch (direction) {
      case 'r':
        if (yAxisInsideBorder(y) && (x >= n || isAlreadyWentThrough(x, y))) {
          x--;
          direction = 'd'
          moveByDir[direction]()
        }
        break;
      case 'd':
        if ((y >= n || isAlreadyWentThrough(x, y)) && xAxisInsideBorder(x)) {
          y--;
          direction = 'l'
          moveByDir[direction]()
        }
        break;
      case 'l':
        if (yAxisInsideBorder(y) && (x < 0 || isAlreadyWentThrough(x, y))) {
          x++;
          direction = 'u'
          moveByDir[direction]()
        }
        break;
      case 'u':
        if ((y < 0 || isAlreadyWentThrough(x, y)) && xAxisInsideBorder(x)) {
          y++;
          direction = 'r'
          moveByDir[direction]()
        }
        break;
    }

    output[y][x] = i

    moveByDir[direction]()
  }

  return output
};
