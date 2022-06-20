/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */

// reference answer from Vivian Liao
var nextPermutation = function (nums) {
    let successor = nums.length - 1, pivot = -1
    // find the pivot and successor
    for (let i = nums.length - 1; i >= 0; i--) {
        // pivot condition: nums[i] > nums[i - 1]
        if (nums[i] > nums[i - 1]) {
            pivot = i - 1
            // successor condition: for nums[successor] > nums[pivot]
            while (nums[successor] <= nums[pivot]) {
                successor -= 1
            }
            break
        }
    }
    // swap pivot and successor element
    if (pivot !== -1 || successor !== nums.length - 1) {
        [nums[successor], nums[pivot]] = [nums[pivot], nums[successor]]
    }
    // from pivot + 1 to end, reverse the rest elements
    // cannot use slice(), reverse(), concat() cuz we need to do it in place
    for (let i = pivot + 1, j = nums.length - 1; i < nums.length; i++) {
        if (i < j) {
            [nums[i], nums[j]] = [nums[j], nums[i]]
            j--;
        }
    }
};