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

// function sumAllNodes(root) {
//     if(root === null) return 0;
//     return root.val + sumAllNodes(root.left) + sumAllNodes(root.right)
// }

var convertBST = function(root) {   
    let sum = 0;
    function convert(root) {
        if(root === null) return null;
        convert(root.right)
        root.val = sum += root.val;
        convert(root.left);
    }
    // const sum = sumAllNodes(root);
    // function replaceWithNewVal(origin) {
    //     if(origin === null) return null;
    //     const newLeftTree = replaceWithNewVal(origin.left);
    //     const leftSum = sumAllNodes(origin.left);
    //     const newVal = sum - leftSum;
    //     origin.val = newVal 
    //     const newRightTree = replaceWithNewVal(origin.right);
    //     const newVal2 = sum - leftSum - origin.val
    //     origin.val = newVal2
    // }
    // replaceWithNewVal(root);
    convert(root);
    return root;
};