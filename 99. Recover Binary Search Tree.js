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
 * @return {void} Do not return anything, modify root in-place instead.
 */

// failed, correct answer
var recoverTree = function (root) {
  let firstEl = null, secondEl = null, prevEl = new TreeNode(Number.MIN_SAFE_INTEGER);
  function _recursive(root) {
    if (root === null) return
    _recursive(root.left)

    if (firstEl === null && prevEl.val >= root.val) {
      firstEl = prevEl
    }
    if (firstEl !== null && prevEl.val >= root.val) {
      secondEl = root
    }
    prevEl = root

    _recursive(root.right)
  }

  _recursive(root);
  let temp = firstEl.val
  firstEl.val = secondEl.val
  secondEl.val = temp
  return root;
};