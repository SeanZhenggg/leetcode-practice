/**
 * @param {number[]} nums
 * @return {number}
 */


// wrong answer by myself
var findUnsortedSubarray = function(nums) {
    let min = nums[0], max = nums[0]
    let minIndex = 0
    let start = 0, end = 0

    for(let i = 1; i < nums.length; i++) {
        if(nums[i] < min) {
            start = Math.min(start, minIndex)
            min = nums[i]
            minIndex = i
        }
        else if(nums[i] >= min && nums[i] <= max) {
            if(!start) start = i - 1
            end = i;
        }
        else {
            max = nums[i]
        }
    }
    
    console.log('start', start)
    console.log('end', end)
    return !end && !start ? 0 : end - start + 1
};

// correct answer 1
var findUnsortedSubarray = function(nums) {
    let subArrLeft = 0, subArrRight = nums.length - 1
    while(subArrLeft < nums.length && nums[subArrLeft] <= nums[subArrLeft+1]) {
        subArrLeft += 1
    }

    while(subArrRight > 0 && nums[subArrRight] >= nums[subArrRight-1]) {
        subArrRight -= 1
    }
    
    if(subArrLeft > subArrRight) return 0
    // for now, the min sub array that need to reorder will be nums[minSubArrStart:minSubArrEnd]
    
    let start = subArrLeft, end = subArrRight
    let min = Math.min(...nums.filter((_, i) => i >= subArrLeft && i <= subArrRight))
    let max = Math.max(...nums.filter((_, i) => i >= subArrLeft && i <= subArrRight))
    while(start >= 0 || end <= nums.length - 1) {
        if(nums[start] > min) subArrLeft = start
        if(nums[end] < max) subArrRight = end
  
        if(start >= 0) start--;
        if(end <= nums.length - 1) end++;
    }
    return subArrRight - subArrLeft + 1
};

// correct answer 2
var findUnsortedSubarray = function(nums) {
    const last = nums.length - 1
    let min = nums[last], max = nums[0]
    let start = -1, end = -2
    
    for(let i = 1; i <= last; i++) {
        max = Math.max(max, nums[i])
        min = Math.min(min, nums[last - i])
        
        if(nums[i] < max) end = i
        if(nums[last - i] > min) start = last - i
    }

    return end - start + 1
};