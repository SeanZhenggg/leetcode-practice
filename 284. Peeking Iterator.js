/**
 * // This is the Iterator's API interface.
 * // You should not implement it, or speculate about its implementation.
 * function Iterator() {
 *    @ return {number}
 *    this.next = function() { // return the next number of the iterator
 *       ...
 *    }; 
 *
 *    @return {boolean}
 *    this.hasNext = function() { // return true if it still has numbers
 *       ...
 *    };
 * };
 */

// my solution
/**
 * @param {Iterator} iterator
 */
var PeekingIterator = function(iterator) {
    this.iterator = iterator
    this.lastStep = {
        command: '',
        val: -1
    }
};

/**
 * @return {number}
 */
PeekingIterator.prototype.peek = function() {
    console.log('this.lastStep', this.lastStep)
    if(this.lastStep.command === 'peek') {
        return this.lastStep.val
    }
    let next = this.iterator.next();
    this.lastStep.command = 'peek'
    this.lastStep.val = next
    return next;
};

/**
 * @return {number}
 */
PeekingIterator.prototype.next = function() {
    console.log('this.lastStep', this.lastStep)
    if(this.lastStep.command === 'peek') {
        this.lastStep.command = 'next'
        return this.lastStep.val
    }
    let next = this.iterator.next()
    this.lastStep.command = 'next'
    this.lastStep.val = next
    return next
};

/**
 * @return {boolean}
 */
PeekingIterator.prototype.hasNext = function() {
    if(this.lastStep.command === 'peek') {
        return true
    } else {
        return this.iterator.hasNext()   
    }
};

/** 
 * Your PeekingIterator object will be instantiated and called as such:
 * var obj = new PeekingIterator(arr)
 * var param_1 = obj.peek()
 * var param_2 = obj.next()
 * var param_3 = obj.hasNext()
 */

// better solution

class PeekingIterator {
    
  constructor(iterator) {
      this.itr = iterator;
      this.peeked = null;
  }
  
  peek() {
      if(!this.peeked) this.peeked = this.itr.next();
      return this.peeked;
  }
  
  next() {
      if(!this.peeked) return this.itr.next()
      const temp = this.peeked;
      this.peeked = null;
      return temp;
  }
  
  hasNext() {
      if(!this.peeked) return this.itr.hasNext()
      return true;
  }
}