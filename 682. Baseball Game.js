/**
 * @param {string[]} ops
 * @return {number}
 */
 var calPoints = function(ops) {
  var records = []
  for(let i = 0; i < ops.length; i++) {
      switch(ops[i]) {
          case 'C':
              records.pop()
              break
          case 'D':
              records.push(records[records.length - 1] * 2)
              break
          case '+':
              records.push(records[records.length - 1] + records[records.length - 2])
              break
          default:
              records.push(Number(ops[i]))
              break
      }
  }
  return records.reduce((sum, val) => sum + val, 0)
};