/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
 var maxOperations = function(nums, k) {
  let hashMap = new Map() // use map to record [k - x]: count
  let result = 0
  for(let i = 0; i < nums.length; i++) {
      var val = hashMap.get(nums[i])
      // if nums[i] === k - nums[i] and map[nums[i]] > 0
      // has stored the other value of pairs before
      if(val && val > 0) {
          result++;
          hashMap.set(nums[i], --val) // decrease count of nums[i]
      // else, no match
      // record value: k - nums[i] into map for later use
      } else {
          var val = hashMap.get(k - nums[i]) || 0
          hashMap.set(k - nums[i], ++val) // increase count of k - nums[i]
      }
  }
  return result.length
};