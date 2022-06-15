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
 * @return {TreeNode}
 */

var increasingBST = function(root) {
    let newRoot = null, lastNode = null;
    function createTree(root) {
        if(root === null) return null;
        createTree(root.left);
        if(!newRoot) {
            newRoot = root
            lastNode = newRoot
        } else {
            lastNode.right = root
            lastNode = lastNode.right
            lastNode.left = null
        }
        createTree(root.right);
    }
    createTree(root);
    return newRoot;
};