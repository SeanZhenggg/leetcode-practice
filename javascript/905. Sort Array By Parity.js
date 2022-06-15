/**
 * @param {number[]} A
 * @return {number[]}
 */
 var sortArrayByParity = function(A) {
  let indexArray = new Array(Math.max(...A) + 1).fill(0)
  let output = []
  for(let i = 0;i < A.length; i++) indexArray[A[i]] ++
  for(let i = 0; i < indexArray.length; i++) {
      for(let j = 0; j < indexArray[i]; j++) {
      if(i % 2 === 0) output.unshift(i)
      else output.push(i)
      }
  }
  return output
};

var sortArrayByParity = function(nums) {
  return nums.reduce((result, num) => {
      if(num % 2 === 0) result.unshift(num)
      else result.push(num)
      return result
  }, [])
};
/**
 * @param {number[]} nums
 * @return {number[]}
 */
 var sortArrayByParity = function(nums) {
  let i = 0, j = nums.length - 1
  while(i < j) {
      if(nums[i] % 2 === 0) i++;
      else if (nums[j] % 2 === 1) j--;
      else {
          let temp = nums[i]
          nums[i] = nums[j]
          nums[j] = temp
          i++;
          j--;
      }
  }
  return nums
};