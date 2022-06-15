/**
 * @param {number[][]} points
 * @return {number}
 */

// answer from MST: Kruskal's Algorithm
// http://alrightchiu.github.io/SecondRound/minimum-spanning-treekruskals-algorithm.html
 var minCostConnectPoints = function(points) {
  // Set data structure methods
  function _findSetAndCollapse(subset, i) {
      let root = i
      while(subset[root] >= 0) {
          root = subset[root]
      }

      // collapse
      while(i !== root) {
          let parent = subset[i]
          subset[i] = root
          i = parent
      }
      return root
  }
  function _UnionSet(subset, x, y) {
      let xRoot = _findSetAndCollapse(subset, x);
      let yRoot = _findSetAndCollapse(subset, y);
      
      if(subset[xRoot] <= subset[yRoot]) {
          subset[xRoot] += subset[yRoot]
          subset[yRoot] = xRoot
      } else {
          subset[yRoot] += subset[xRoot]
          subset[xRoot] = yRoot
      }
  }
  // variables
  let vertex = {}
  let allEdges = Array.from({ length: points.length }).map(_ => Array.from({ length: points.length }).fill(0))
  let increaseWeight = []
  let edgesetMST = []
  let subset = Array.from({ length: points.length }).fill(-1)
  // initial vertices map
  for(let i = 0; i < points.length; i++) {
      vertex[i] = { x: points[i][0], y: points[i][1] }
  }
  // initial all edges coord, weight
  for(let i = 0; i < points.length; i++) {
      for(let j = 0; j < points.length; j++) {
          allEdges[i][j] = Math.abs(vertex[i].x - vertex[j].x) + Math.abs(vertex[i].y - vertex[j].y) 
      }
  }
  // initial increase weight(edges used and sorted by ascending order)
  for(let i = 0; i < points.length; i++) {
      for(let j = i+1; j < points.length; j++) {
          increaseWeight.push({ from: i, to: j, w: allEdges[i][j] })
      }
  }
  increaseWeight = increaseWeight.sort((a, b) => a.w - b.w)
  // Kruskal's Algorithm
  for(let i = 0; i < increaseWeight.length; i++) {
      const { from, to, w} = increaseWeight[i]
      if(_findSetAndCollapse(subset, from) !== _findSetAndCollapse(subset, to)) {
          edgesetMST.push({ from: from, to: to, w: w })
          _UnionSet(subset, from, to)
      }
  }

  return edgesetMST.reduce((result, edge) => result += edge.w, 0)
};