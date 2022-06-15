/**
 * @param {string} s
 * @return {boolean}
 */
 var validPalindrome = function(s) {
  var i = 0;
  var j = s.length - 1;
  function _recursive(i, j, isDelete) {
      while(i < j) {
          if(s[i] !== s[j]) {
              if(isDelete) {
                  return false;
              }
              return _recursive(i + 1, j, true) || _recursive(i, j - 1, true)
          }
          i ++;
          j --;
      }
      return true;
  }
  
  return _recursive(i, j, false)
};