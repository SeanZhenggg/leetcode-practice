/**
 * @param {number[]} nums
 * @return {boolean}
 */
var find132pattern = function(nums) {
   var monotonicSt = [], s3 = Number.MIN_SAFE_INTEGER
   for(let i = nums.length - 1; i >= 0; i--) {
       if(nums[i] < s3) return true
       while(monotonicSt.length > 0 && monotonicSt[monotonicSt.length - 1] < nums[i]) {
           s3 = monotonicSt.pop()
       }
       monotonicSt.push(nums[i])
   }
    
   return false;
};