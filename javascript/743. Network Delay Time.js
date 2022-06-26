/**
 * @param {number[][]} times
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
/**
 * @param {number[][]} times
 * @param {number} n
 * @param {number} k
 * @return {number}
 */

// min heap answer
class HeapNode {
  constructor(element, key) {
      this.element = element
      this.key = key
  }
}
class Heap {
  _heap;
  constructor() {
    this._heap = []
  }
  swap(a, b) { [this._heap[a], this._heap[b]] = [this._heap[b], this._heap[a]] }
  top = () => this._heap[0]
  isHeapEmpty = () => this.size() < 1
  size = () => this._heap.length
  getHeap = () => this._heap
  getParentNode = node => Math.floor((node - 1) / 2)
  findPosition = element => this._heap.findIndex(node => node.element === element)
  heapify(node, length) {}
  change(element, newKey) {}
  extract() {}
}
class MinHeap extends Heap {
  constructor() {
    super()
  }
  heapify(node, length) {
    let left = (2 * node) + 1, right = (2 * node) + 2, smallest = node

    if(left <= length && this._heap[left].key < this._heap[node].key) {
      smallest = left
    }
    if(right <= length && this._heap[right].key < this._heap[smallest].key) {
      smallest = right
    }

    if(smallest !== node) {
      this.swap(smallest, node)
      this.heapify(smallest, length)
    }
  }
  change(element, newKey) {
    if(this.isHeapEmpty()) {
      console.log('heap is empty')
      return false
    }
    let index_node = this.findPosition(element)
    if(this._heap[index_node].key < newKey) {
      console.log('new key is greater than current key')
      return false;
    }

    this._heap[index_node].key = newKey

    while (index_node > 0 && this._heap[this.getParentNode(index_node)].key >= this._heap[index_node].key) {
      this.swap(index_node, this.getParentNode(index_node))
      index_node = this.getParentNode(index_node)
    }
    return true;
  }
  insert(element, key) {
    if(this.findPosition(element) === -1) {
        this._heap.push(new HeapNode(element, key))  
    }

    const result = this.change(element, key)
    return result
  }
  extract() {
    if(this.isHeapEmpty()) {
      console.log('heap is empty')
      return null
    }
    let min = this.top()
    this._heap[0] = this._heap[this.size() - 1]
    this._heap.pop()
    this.heapify(0, this.size() - 1)
    return min
  }
}
var networkDelayTime = function(times, n, k) {
  // step 1, 2
  const edges = new Map()
  const visit = new Set()
  for (let [u, v, w] of times) {
      if(!edges.has(u)) edges.set(u, [{ v, w }])
      else edges.set(u, [...edges.get(u), { v, w }])
  }
  // step 3
  const minHeap = new MinHeap()
  minHeap.insert(k, 0)

  let ret = -Infinity
  while(!minHeap.isHeapEmpty()) { // step 4
      // step 5
      let { element, key: weight } = minHeap.extract() // get path of smallest weight, with its target(current) vertex and weight
      if(visit.has(element)) continue // continue if visit has element included
      visit.add(element) // add current vertex into visit

      ret = Math.max(ret, weight) // step 6
      
      const edgesForEle = edges.get(element) || [] // get all paths with current vertex from edge map
      for (let { v, w } of edgesForEle) {
          if(v && (w || w === 0) && !visit.has(v)) { // if not visited yet
              const result = minHeap.insert(v, weight + w) // add the target vertex and weight into heap, weight calculated with current vertex weight 
          }
      }
  }
  
  return visit.size === n ? ret : -1
  
};

// dictionary answer
var networkDelayTime = function (times, n, k) {
  var visited = []
  var dictionary = {}
  var edges = new Map()

  for (let [u, v, w] of times) {
    if (!edges.has(u)) edges.set(u, [{ v, w }])
    else {
      edges.set(u, [...edges.get(u), { v, w }])
    }
  }

  for (let i = 1; i <= n; i++) {
    dictionary[i] = {
      shortestPath: Infinity,
      previousVertex: null
    }
  }

  dictionary[k].shortestPath = 0

  while (visited.length !== n) {
    let currentVertex = Object.keys(dictionary).filter(e => !visited.includes(e)).sort((a, b) => dictionary[a].shortestPath - dictionary[b].shortestPath)[0]
    let allEdges = edges.get(Number(currentVertex))
    if (allEdges) {
      for (let edge of allEdges) {
        if (!visited.includes(edge.v)) {
          if (edge.w + dictionary[currentVertex].shortestPath < dictionary[edge.v].shortestPath) {
            dictionary[edge.v].shortestPath = edge.w + dictionary[currentVertex].shortestPath
            dictionary[edge.v].previousVertex = currentVertex
          }
        }
      }
    }
    visited.push(currentVertex)
  }

  if (Object.keys(dictionary).findIndex(key => dictionary[key].shortestPath === Infinity) !== -1) {
    return -1;
  } else {
    return Math.max(...Object.keys(dictionary).map(key => dictionary[key].shortestPath))
  }
};