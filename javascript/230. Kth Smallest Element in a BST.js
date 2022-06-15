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
 * @param {number} k
 * @return {number}
 */
 var kthSmallest = function(root, k) {
    let smallest, dk = k;
    function findSmallest(root) {
        if(root === null) return;
        findSmallest(root.left)
        if(smallest || smallest === 0) return;
        if(dk > 0) dk--;
        if(dk === 0) {
            smallest = root.val
            return 
        }
        findSmallest(root.right)
    }
    findSmallest(root);
    return smallest
};