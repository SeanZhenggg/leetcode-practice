/**
 * @param {number[]} nums
 * @return {number[][]}
 */
// my answer
var permuteUnique = function (nums) {
  let resultMap = {}
  let result = []
  function _backtracking(rest, combination = []) {
    if (!rest.length) {
      let comKey = combination.join(',')
      if (!resultMap[comKey]) {
        resultMap[comKey] = true
        result.push(combination)
      }
      return
    }
    for (let i = 0; i < rest.length; i++) {
      let newCombination = combination.concat(rest[i]).slice()
      let newRest = rest.slice()
      newRest.splice(i, 1)
      _backtracking(newRest, newCombination)
    }
  }

  _backtracking(nums)
  return result
};

// reference answer
var permuteUnique = function(nums) {
  const sorted = nums.sort((x,y) => x-y), permutations = [];

  const rcr = (arr, permutation) => {
      if (!arr.length) return permutations.push(permutation);

      let prev = -Infinity;
      for (let i = 0; i < arr.length; i++) {
          if (prev === arr[i]) continue;

          newArr = arr.slice(0, i).concat(arr.slice(i+1));
          rcr(newArr, [...permutation, arr[i]]);

          prev = arr[i];
      }
  }
  rcr(nums, []);

  return permutations;
};

// reference answer 2
var permuteUnique = function(nums) {
  let res = [];
  
  const dfs = (nums, i) => {
      console.log('nums = ', nums)
      //base 
      if(i === nums.length){
          console.log('push!!!!!!!!!')
          res.push(nums.slice());
          return
      }
      let hash = {};
      //loop for changing the other elements
      for(let j = i; j<nums.length; j++){
          if(hash[nums[j]]) continue;
          hash[nums[j]] = true;
          [nums[i], nums[j]] = [nums[j], nums[i]];
          console.log(`i = ${i}`, `j = ${j}`)
          console.log(`before: switch nums[i] ${nums[i]} and nums[j] ${nums[j]} `)
          dfs(nums, i+1);
          [nums[i], nums[j]] = [nums[j], nums[i]];
          console.log(`after: switch nums[j] ${nums[j]} and nums[i] ${nums[i]} `)
      }
  }
  
  dfs(nums, 0);
  return res;
};


permuteUnique([1,2,3])