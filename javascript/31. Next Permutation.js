/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */

// reference answer from Vivian Liao
 var nextPermutation = function(nums) {
  let successor = nums.length - 1, pivot = -1
  
  for(let i = nums.length - 1; i >= 0; i--) {
      if(nums[i] > nums[i - 1]) {
          pivot = i - 1
          while(successor > pivot && nums[successor] <= nums[pivot]) {
              successor -= 1
          }
          break
      }
  }

  if(pivot !== -1 || successor !== nums.length - 1) {
      [nums[successor], nums[pivot]] = [nums[pivot], nums[successor]]
  }
  for(let i = pivot + 1, j = nums.length - 1; i < nums.length; i++) {
      if(i < j) {
          [nums[i], nums[j]] = [nums[j], nums[i]]
          j --;
      }
  }
};