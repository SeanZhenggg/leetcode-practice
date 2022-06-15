/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */

// failed answer
 var min = { val: null }
 var max = { val: null }
 var isValidBST = function(root) {
     let isValid = true;
     let defaultMax = Number.MAX_SAFE_INTEGER
     let defaultMin = Number.MIN_SAFE_INTEGER
     min.val = null
     max.val = null
     function _recursive(root) {
         console.log('-----start------')
         if(root === null) return true;
         console.log('root.val is at', root.val)
         const rightIsValid = _recursive(root.right);
         if(!rightIsValid || (root.right && root.val >= root.right.val) || (min.val && root.val >= min.val)) return false
         const leftIsValid = _recursive(root.left);
         if(!leftIsValid || (root.left && root.val <= root.left.val) || (max.val && root.val <= max.val)) return false
         
         let lastMin = min.val === null ? defaultMax : min.val
         let leftVal = root.left?.val === undefined ? defaultMax : root.left.val
         let rightVal = root.right?.val === undefined ? defaultMax : root.right.val
         min.val = Math.min(lastMin, root.val, leftVal, rightVal)
         let lastMax = max.val === null ? defaultMin : max.val
         leftVal = root.left?.val === undefined ? defaultMin : root.left.val
         rightVal = root.right?.val === undefined ? defaultMin : root.right.val
         max.val = Math.max(lastMax, root.val, leftVal, rightVal)
         console.log('min', min)
         console.log('max', max)
         console.log('leftIsValid', leftIsValid, 'rightIsValid', rightIsValid)
         console.log(`-----end from root.val = ${root.val}------`)
         return leftIsValid && rightIsValid
     }
     
     return _recursive(root)
 };


// correct answer
var isValidBST = function(root) {
  let minSatisfiedValue = Number.MIN_SAFE_INTEGER
  let isValid = true
  function _recursive(root) {
      if(!root) return
      _recursive(root.left)
      if(root.val > minSatisfiedValue) {
          minSatisfiedValue = root.val
      } else {
          isValid = false
      }
      _recursive(root.right)
  }
  
  _recursive(root)
  return isValid;
};

// referenced answer
var isValidBST = function(root, min = Number.MIN_SAFE_INTEGER, max = Number.MAX_SAFE_INTEGER) {
  if(!root) return true
  if(root.val < min || root.val > max) return false

  return isValidBST(root.left, min, root.val) && isValidBST(root.right, root.val, max)
}