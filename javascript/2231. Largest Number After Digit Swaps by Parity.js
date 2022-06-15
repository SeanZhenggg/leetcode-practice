/**
 * @param {number} num
 * @return {number}
 */
// var largestInteger = function(num) {
//     let odd = [], oddPosition = {}, even = [], evenPosition = {}
//     String(num).split('').forEach((e, index) => {
//         if(e%2 === 0) {
//             even.push(e)
//             evenPosition[index] = true
//         }
//         else {
//             odd.push(e)
//             oddPosition[index] = true
//         }
//     })
    
//     var newEven = even.sort()
//     var newOdd = odd.sort()
//     var i = 0, j = 0
//     return Array.from({ length: num.length }).reduce((result, e, index) => {
//         if(evenPosition[index]) {
//             result.push(Number(newEven[i]))
//             i++
//         }
//         if(oddPosition[index]) {
//             result.push(Number(newOdd[j]))
//             j ++
//         }
//     }, [])
// };

var largestInteger = function(num) {
    let odd = [], oddPosition = {}, even = [], evenPosition = {}
    String(num).split('').forEach((e, index) => {
        if(e%2 === 0) {
            even.push(e)
            evenPosition[index] = true
        }
        else {
            odd.push(e)
            oddPosition[index] = true
        }
    })
    
    var newEven = even.sort((a, b) => b.localeCompare(a))
    var newOdd = odd.sort((a, b) => b.localeCompare(a))
    var i = 0, j = 0
    return Array.from({ length: String(num).length }).reduce((result, e, index) => {
        if(evenPosition[index]) {
            result.push(Number(newEven[i]))
            i++
        }
        if(oddPosition[index]) {
            result.push(Number(newOdd[j]))
            j ++
        }
        return result
    }, []).join('')
};