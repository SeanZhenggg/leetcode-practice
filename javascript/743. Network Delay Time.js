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
    let = currentVertex = Object.keys(dictionary).filter(e => !visited.includes(e)).sort((a, b) => dictionary[a].shortestPath - dictionary[b].shortestPath )[0]
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

  if(Object.keys(dictionary).findIndex(key => dictionary[key].shortestPath === Infinity) !== -1) {
    return -1;
  } else {
    return Math.max(...Object.keys(dictionary).map(key => dictionary[key].shortestPath))
  }
};