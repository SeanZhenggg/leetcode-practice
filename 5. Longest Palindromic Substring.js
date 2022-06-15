// https://leetcode.com/problems/longest-palindromic-substring/

/**
 * @param {string} s
 * @return {string}
 */

// first answer
var longestPalindrome = function(s) {
    if(s.length === 1) return s;
    let ans = ''
    function _isPalindrome(substr) {
        let l = 0, r = substr.length - 1
        while (l < r) {
            if(substr[l] !== substr[r]) return false
            l++;
            r--;
        }
        return true;
    }
    
    for(let i = 0; i < s.length - 1; i++) {
        let r = i;
        while (r <= s.length - 1) {
            let substr;
            if(r === s.length - 1) {
                substr = s.substring(i);
            } else {
                substr = s.substring(i, r+1)
            }
            if(_isPalindrome(substr)) {
                if(ans.length < substr.length) ans = substr;
            }
            r++;
        }
    }
    
    return ans;
};

// dp answer
var longestPalindrome = function(s) {
  if(s.length === 1) return s;
  let ans = ''
  let dpTable = Array.from({ length: s.length }).map(_ => Array.from({ length: s.length }).fill(false))
  
  for(let i = s.length - 1; i >= 0; i--) {
    for(let j = i; j < s.length; j++) {
        // console.log('i', i, 'j', j)
      dpTable[i][j] = (s[i] === s[j]) && (j - i < 3 || dpTable[i+1][j-1])
      if(dpTable[i][j] && s.substring(i, j+1).length > ans.length) {
        ans = s.substring(i, j+1)
      }
    }
  }
  return ans;
};
