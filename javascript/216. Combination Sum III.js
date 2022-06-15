/**
 * @param {number} k
 * @param {number} n
 * @return {number[][]}
 */

// first answer
var combinationSum3 = function (k, n) {
  let result = {}
  if (n < k) return []

  function _backtracking(curSum, ele, count) {
    if (count === k) {
      if (curSum === n) {
        var sorted = ele.split('').sort((a, b) => a - b).join('')
        if (!result[sorted]) result[sorted] = true
      }
    }
    else {
      for (let i = 1; i <= 9; i++) {
        if (ele.includes(i)) continue
        _backtracking(curSum + i, ele + i, count + 1)
      }
    }
  }

  _backtracking(0, "", 0)

  return Object.keys(result).map((key) => key.split(''))
};

console.time("first answer")
combinationSum3(5, 20)
console.timeEnd("first answer")

// optimized answer
var optimized_combinationSum3 = function (k, n) {
  let result = []
  if (n < k) return []

  function _backtracking(curSum, ele, start) {
    if (curSum > n) return
    if (ele.length === k) {
      if (curSum === n) result.push(ele)
      return
    }
    for (let i = start; i <= 9; i++) {
      _backtracking(curSum + i, [...ele, i], i + 1)
    }
  }

  _backtracking(0, [], 1)

  return result
};

console.time("optimized answer")
optimized_combinationSum3(5, 20)
console.timeEnd("optimized answer")