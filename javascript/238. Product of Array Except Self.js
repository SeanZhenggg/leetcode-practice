/**
 * @param {number[]} nums
 * @return {number[]}
 */

// using left, right records
var productExceptSelf = function(nums) {
    let size = nums.length
    let left = Array.from({ length: size })
    let right = Array.from({ length: size })
    let ans = Array.from({ length: size })

    left[0] = 1
    right[size - 1] = 1
    for(let lp = 1; lp <= size - 1; lp++) {
        left[lp] = left[lp - 1] * nums[lp - 1]
    }
    for(let rp = size - 2; rp >= 0; rp--) {
        right[rp] = right[rp + 1] * nums[rp + 1]
    }

    for(let i = 0; i <= size - 1; i++) {
        ans[i] = left[i] * right[i]
    }
    return ans
};

// without extra space except ans array
var productExceptSelf = function (nums) {
  let size = nums.length
  let ans = []
  ans[0] = 1
  let rightProd = 1

  for (let lp = 1; lp <= size - 1; lp++) {
    ans[lp] = ans[lp - 1] * nums[lp - 1]
  }
  for (let rp = size - 1; rp >= 0; rp--) {
    ans[rp] = ans[rp] * rightProd
    rightProd *= nums[rp]
  }

  return ans
};

// without extra space except ans array and using only one loop
// with two pointer and two variables for recording prefix product and postfix product
var productExceptSelf = function(nums) {
  let last = nums.length - 1
  let LPtr = 0, RPtr = last
  let LPreProd = 1, RPosProd = 1
  let result = new Array(nums.length).fill(1)

  while (LPtr <= last && RPtr >= 0) {
      result[LPtr] *= LPreProd
      LPreProd *= nums[LPtr]
      
      result[RPtr] *= RPosProd
      RPosProd *= nums[RPtr]
      
      LPtr ++;
      RPtr --;
  }

  return result
};

console.log(productExceptSelf(nums))