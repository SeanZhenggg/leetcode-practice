/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */

// array answer
var topKFrequent = function (nums, k) {
  if (nums.length === 0) return []


  var map = nums.reduce((map, val) => {
    map[val] = map[val] || 0
    map[val]++
    return map
  }, {})

  return Object.keys(map).sort((a, b) => map[b] - map[a]).slice(0, k)
};

// maxHeap answer
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */

class HeapNode {
  element;
  weight;
  constructor(ele, w) {
    this.element = ele
    this.weight = w
  }
}

class MaxHeap {
  heap = []
  constructor(elements) {
    elements.forEach(element => {
      var node = new HeapNode(element.k, element.w)
      this.heap.push(node)
    });
    for (let i = Math.floor((elements.length - 1) / 2); i >= 0; i--) {
      this.heapify(i, elements.length - 1)
    }
  }
  swap(a, b) {
    [this.heap[a], this.heap[b]] = [this.heap[b], this.heap[a]]
  }
  heapify(node, length) {
    let left = node * 2 + 1, right = node * 2 + 2, biggest = node

    if (left <= length && this.heap[node].weight < this.heap[left].weight) {
      biggest = left
    }

    if (right <= length && this.heap[biggest].weight < this.heap[right].weight) {
      biggest = right
    }

    if (biggest !== node) {
      this.swap(biggest, node)
      this.heapify(biggest, length)
    }
  }
  extract() {
    if (this.heap.length === 0) return
    let min = this.heap[0]
    this.heap[0] = this.heap.splice(-1, 1)[0]
    this.heapify(0, this.heap.length - 1)
    return min
  }
}

var topKFrequent = function (nums, k) {
  let result = []
  let map = nums.reduce((map, num) => {
    map[num] = (map[num] + 1) || 1
    return map
  }, {})
  let mapArr = Object.keys(map).map(key => ({
    k: Number(key),
    w: map[key]
  }))

  const maxHeap = new MaxHeap(mapArr)

  while (k-- > 0) {
    const node = maxHeap.extract()
    result.push(node.element)
  }

  return result
};