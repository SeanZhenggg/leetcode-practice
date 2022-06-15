/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  let i = 0
  let matched = {
      '(': ')',
      '[': ']',
      '{': '}',
  }
  let left = []
  if(s.length === 1) return false;

  while(i <= s.length - 1) {
      if(s[i] === '(' || s[i] === '[' || s[i] === '{') {
          left.push(s[i])
      }
      if(s[i] === ')' || s[i] === ']' || s[i] === '}') {
          const lastParentheses = left.pop();
          if(matched[lastParentheses] !== s[i]) return false;
      }
      i++;
  }
  if(left.length > 0) return false;
  return true;
};
