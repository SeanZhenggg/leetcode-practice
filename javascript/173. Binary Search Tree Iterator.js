/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

//#region first answer
/**
 * @param {TreeNode} root
 */
 var BSTIterator = function(root) {
  this.treeList = []
  function goThrough(root) {
      if(root === null) return root
      goThrough.call(this, root.left)
      this.treeList.push(root.val)
      goThrough.call(this, root.right)
      return;
  }
  goThrough.call(this, root);
  this.pointer = null
  return this
};

/**
* @return {number}
*/
BSTIterator.prototype.next = function() {
  if(this.pointer === null) {
      this.pointer = 0
      return this.treeList[this.pointer];
  }
  return this.treeList[++this.pointer];
};

/**
* @return {boolean}
*/
BSTIterator.prototype.hasNext = function() {
  if(this.pointer === null) {
      return true;
  }
  let pointer = this.pointer
  if(this.treeList[++pointer]) return true;
  return false;
};
//#endregion

//#region stack answer
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
 */
 var BSTIterator = function(root) {
  this.treeStack = []
  this.goThrough = function(root) {
      if(root === null) return root
      this.treeStack.unshift(root)
      this.goThrough(root.left)
      return;
  }
  this.goThrough(root);
  return this
};

/**
* @return {number}
*/
BSTIterator.prototype.next = function() {
  const stackTop = this.treeStack.shift();
  if(stackTop.right !== null) {
      this.goThrough(stackTop.right)
  }
  return stackTop.val;
};

/**
* @return {boolean}
*/
BSTIterator.prototype.hasNext = function() {
  return this.treeStack.length > 0
};

/** 
* Your BSTIterator object will be instantiated and called as such:
* var obj = new BSTIterator(root)
* var param_1 = obj.next()
* var param_2 = obj.hasNext()
*/
/** 
* Your BSTIterator object will be instantiated and called as such:
* var obj = new BSTIterator(root)
* var param_1 = obj.next()
* var param_2 = obj.hasNext()
*/