/**
 * @param {number} n
 * @return {number[][]}
 */
var generateMatrix = function (n) {
  let output = Array.from({ length: n }).map(e => Array.from({ length: n }))
  let x = 0, y = 0
  let dir = 'r'

  let _getNewCords = () => {
    switch (dir) {
      case 'r':
        y++;
        break;
      case 'd':
        x++;
        break;
      case 'l':
        y--;
        break;
      case 'u':
        x--;
        break;
      default:
        break;
    }
  }
  for (let i = 1; i <= n * n; i++) {
    if (i === 1) {
      output[x][y] = i
      _getNewCords()
      continue
    }

    if (dir === 'r' && x >= 0 && x <= n - 1 && (y >= n || output[x][y] !== undefined)) {
      dir = 'd'
      y--;
      x++;
    }
    if (dir === 'd' && (x >= n || output[x][y] !== undefined) && y >= 0 && y <= n - 1) {
      dir = 'l'
      x--;
      y--;
    }
    if (dir === 'l' && x >= 0 && x <= n - 1 && (y < 0 || output[x][y] !== undefined)) {
      dir = 'u'
      y++;
      x--;
    }
    if (dir === 'u' && (x < 0 || output[x][y] !== undefined) && y >= 0 && y <= n - 1) {
      dir = 'r'
      x++;
      y++;
    }

    output[x][y] = i

    _getNewCords()
  }

  return output
};