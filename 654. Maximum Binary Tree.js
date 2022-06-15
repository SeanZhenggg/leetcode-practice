var constructMaximumBinaryTree = function(nums) {
  let result = createTree(nums);
  return result;
};

function createTree(numbers) {
  if(numbers.length === 0) return null
  var maxInfo = findMax(numbers)
  var subArrPrefix = numbers.slice(0, maxInfo.index)
  var subArrPostfix = numbers.slice(maxInfo.index + 1)
  const node = new TreeNode(maxInfo.val, createTree(subArrPrefix), createTree(subArrPostfix));
  return node
}

function findMax(nums) {
  return nums.reduce(
      (max, num, index) => max.val < num ? { val: num, index } : max 
      , { val: -1, index: -1 }
  )
}