/**
 * @param {string} s
 * @param {number[][]} pairs
 * @return {string}
 */

// answer from union find
var smallestStringWithSwaps = function(s, pairs) {
  // variables
  let subset = Array.from({ length: s.length }).fill(-1)
  let map = {}
  let result = ''
  // union-find methods
  function _find(subset, i) {
      let root = i
      for (root = i; subset[root] >= 0; root = subset[root]);
      return root
  }  
  function _union(subset, x, y) {
      let xRoot = _find(subset, x)
      let yRoot = _find(subset, y)
      // important!!! when x root === y root
      if(xRoot === yRoot) return;
      if(subset[xRoot] <= subset[yRoot]) {
          subset[xRoot] += subset[yRoot];
          subset[yRoot] = xRoot
      } else {
          subset[yRoot] += subset[xRoot];
          subset[xRoot] = yRoot
      }
  }
  // union each pair in pairs
  for(let i = 0; i <pairs.length; i++) {
      _union(subset, pairs[i][0], pairs[i][1])
  }
  // create map for each connected component
  for(let i = 0; i < subset.length; i++) {
      let parent = _find(subset, i)
      map[parent] = map[parent] || []
      map[parent].push(i)
  }
  // sort each connected component in descending order
  for(let key in map) {
      map[key].sort((a, b) => s[b].localeCompare(s[a]))
  }
  // for each character, find its connected component and pop the smallest character
  for(let i = 0; i < s.length; i++) {
    let root = _find(subset, i)
    result += s[map[root].pop()]
  }
  return result;
};