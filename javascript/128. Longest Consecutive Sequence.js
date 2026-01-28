/**
 * @param {number[]} nums
 * @return {number}
 */
var longestConsecutive = function (nums) {
  if (nums.length === 0) return 0
  var sorted = Array.from(new Set(nums.sort((a, b) => a - b)))
  let ret = 1
  let start = 0,
    end = 0
  for (let i = 1; i < sorted.length; i++) {
    if (sorted[i] === sorted[i - 1] + 1) {
      end = i
    } else {
      start = i
      end = i
    }
    ret = Math.max(ret, end - start + 1)
  }
  return ret
}

// union find ans
class UnionFind {
  constructor(nums, parentArray) {
    this.nums = nums
    this.parentArray = parentArray
  }

  union(from, to) {
    const fromParent = this.findParent(from)
    const toParent = this.findParent(to)

    fromParent.parent = toParent.parent
    toParent.count += fromParent.count
  }

  findParent(index) {
    if (this.parentArray[index].parent === index) return this.parentArray[index]

    const parent = this.findParent(this.parentArray[index].parent)
    this.parentArray[index].parent = parent.parent

    return parent
  }

  hasSameParents(x, y) {
    return this.findParent(x) === this.findParent(y)
  }
}

var longestConsecutive = function (nums) {
  function _getMaxConsecutiveSequences(unionFind) {
    let max = 1

    unionFind.parentArray.forEach((item, i) => {
      if (item.parent === i) {
        max = Math.max(max, item.count)
      }
    })

    return max
  }
  function _findConsecutiveSequences(nums, unionFind) {
    let seenNumbers = new Map()

    for (let from = 0; from < nums.length; from++) {
      const currentNumber = nums[from]

      if (seenNumbers.has(currentNumber)) continue

      if (seenNumbers.has(currentNumber - 1)) {
        const to = seenNumbers.get(currentNumber - 1)

        if (!unionFind.hasSameParents(from, to)) unionFind.union(from, to)
      }

      if (seenNumbers.has(currentNumber + 1)) {
        const to = seenNumbers.get(currentNumber + 1)

        if (!unionFind.hasSameParents(from, to)) unionFind.union(from, to)
      }

      seenNumbers.set(currentNumber, from)
    }
  }

  if (nums.length === 0) return 0

  let parentArray = nums.map((number, i) => {
    return { parent: i, count: 1 }
  })

  let unionFind = new UnionFind(nums, parentArray)

  _findConsecutiveSequences(nums, unionFind)

  return _getMaxConsecutiveSequences(unionFind)
}
