
/** 
 * Your MyStack object will be instantiated and called as such:
 * var obj = new MyStack()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.empty()
 */

// first answer
var MyStack = function() {
    this.queue = []
    return this
};

/** 
 * @param {number} x
 * @return {void}
 */
MyStack.prototype.push = function(x) {
    this.queue.push(x)
};

/**
 * @return {number}
 */
MyStack.prototype.pop = function() {
    let popItem = this.queue.pop()
    return popItem;
};

/**
 * @return {number}
 */
MyStack.prototype.top = function() {
    return this.queue[this.queue.length - 1]
};

/**
 * @return {boolean}
 */
MyStack.prototype.empty = function() {
    return this.queue.length === 0
};

// using two queues

var MyStack = function() {
  this.inQueue = []
  this.outQueue = []
  return this
};

/** 
* @param {number} x
* @return {void}
*/
MyStack.prototype.push = function(x) {
  this.inQueue.push(x)
};

/**
* @return {number}
*/
MyStack.prototype.pop = function() {
  while(this.inQueue.length > 1) {
      this.outQueue.push(this.inQueue.shift())
  }
  
  const ele = this.inQueue.shift()
  this.inQueue = this.outQueue
  this.outQueue = []
  return ele
};

/**
* @return {number}
*/
MyStack.prototype.top = function() {
  while(this.inQueue.length > 1) {
      this.outQueue.push(this.inQueue.shift())
  }
  
  const ele = this.inQueue.shift()
  this.outQueue.push(ele)
  this.inQueue = this.outQueue
  this.outQueue = []
  return ele
};

/**
* @return {boolean}
*/
MyStack.prototype.empty = function() {
  return this.inQueue.length === 0
};
