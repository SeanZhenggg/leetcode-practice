/**
 * @param {number} n
 * @return {number}
 */
// first answer - backtracking
// var countVowelStrings = function (n) {
//   let result = []
//   let vowels = ['a', 'e', 'i', 'o', 'u']
//   function _backtracking(chars, index) {
//     if (chars.length === n) {
//       result.push(chars)
//       return
//     }

//     for (let i = index; i < vowels.length; i++) {
//       _backtracking(chars + vowels[i], i)
//     }
//   }

//   _backtracking('', 0)

//   return result.length
// };

/**
 * @param {number} n
 * @return {number}
 */
// advanced answer - dynamic programming
var countVowelStrings = function (n) {
  let dp = Array.from({ length: 5 }).fill(1)

  for (let i = 2; i <= n; i++) {
    for (let j = 3; j >= 0; j--) {
      dp[j] += dp[j + 1]
    }
  }

  return dp.reduce((acc, val) => acc += val, 0)
};

for(let i = 1; i < 10; i++) {
  console.log(countVowelStrings(i))
}