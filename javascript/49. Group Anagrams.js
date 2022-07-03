/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function(strs) {
    if(strs.length === 1) return [strs]
    
    var sortedStrsMap = strs.reduce((map, str, idx) => {
        let sorted = str.split('').sort().join('')
        map[sorted] = map[sorted] || { val: [] }
        map[sorted].val.push(idx)
        return map
    }, {})
    
    return Object.keys(sortedStrsMap).map((key) => {
        return sortedStrsMap[key].val.map(idx => strs[idx])
    })
};