/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    var distinct = new Set();
    nums.forEach((e) => { distinct.add(e) })
    
    return distinct.size !== nums.length
};