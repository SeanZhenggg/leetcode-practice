/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
 var topKFrequent = function(nums, k) {
  if(nums.length === 0) return []

  
  var map = nums.reduce((map, val) => {
      map[val] = map[val] || 0
      map[val] ++
      return map
  }, {})
  
  return Object.keys(map).sort((a, b) => map[b] - map[a]).slice(0, k)
};