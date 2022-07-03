/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */

// sort answer
var isAnagram = function (s, t) {
  return s.split('').sort().join('') === t.split('').sort().join('')
};

// hash table answer
var isAnagram = function (s, t) {
  if (s.length != t.length) {
    return false
  }

  var hashS = {}
  for (let i = 0; i < s.length; i++) {
    hashS[s[i]] = (hashS[s[i]] + 1) || 1
  }

  for (let i = 0; i < t.length; i++) {
    if (!hashS[t[i]]) return false

    hashS[t[i]]--;

    if (hashS[t[i]] < 0) return false
  }
  return true
};