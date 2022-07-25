/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    let left = 0, right = numbers.length - 1
    while (numbers[left] + numbers[right] !== target) {
        let resdualRight = target - numbers[left], resdualLeft = target - numbers[right]
        if(numbers[right] > resdualRight) {
            right --;
        }
        else if(numbers[left] < resdualLeft) {
            left ++;
        }
    }
    return [left + 1, right + 1]
};

var twoSum = function(numbers, target) {
  let left = 0, right = numbers.length - 1
  while (numbers[left] + numbers[right] !== target) {
      if(numbers[left] + numbers[right] > target) {
          right --;
      }
      else if(numbers[left] + numbers[right] < target) {
          left ++;
      }
  }
  return [left + 1, right + 1]
};