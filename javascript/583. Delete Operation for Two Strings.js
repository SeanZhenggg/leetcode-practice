/**
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
// reference answer - dp
var minDistance = function (word1, word2) {
  var dp = Array.from({ length: word1.length + 1 }).map(_ => new Array(word2.length + 1).fill(0))
  var m = word1.length, n = word2.length
  for(let i = 1; i <= m; i++) {
    for(let j = 1; j <= n; j++) {
      if(word1[i - 1] === word2[j - 1]) {
        dp[i][j] = 1 + dp[i - 1][j - 1]
      } else {
        dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1])
      }
    }  
  }

  return (word1.length + word2.length) - 2 * dp[m][n]
};