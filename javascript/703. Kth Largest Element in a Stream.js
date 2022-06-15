/**
 * @param {number} k
 * @param {number[]} nums
 */
var KthLargest = function(k, nums) {
    this.kthLargest = k
    this.result = nums
    this.result.sort((a, b) => a - b)
};

/** 
 * @param {number} val
 * @return {number}
 */
KthLargest.prototype.add = function(val) {
    let index = this.result.findIndex(v => v >= val)
    if(index === -1) this.result.push(val)
    else this.result.splice(index, 0, val)
    var firstIndex = this.result.length - this.kthLargest

    return this.result[firstIndex < 0 ? 0 : firstIndex]
};

/** 
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */