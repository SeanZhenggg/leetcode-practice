/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function (digits) {
    // edge case
    if (digits.length === 0) return []
    // hashMap for number versus letters
    const map = {}
    for (let i = 2, j = 97; i <= 9; i++) {
        if (i === 7 || i === 9) { // number 7, 9 has 4 letters
            map[i] = [String.fromCharCode(j), String.fromCharCode(j + 1), String.fromCharCode(j + 2), String.fromCharCode(j + 3)]
            j += 4
        } else {
            map[i] = [String.fromCharCode(j), String.fromCharCode(j + 1), String.fromCharCode(j + 2)]
            j += 3
        }
    }
    // edge case
    if (digits.length === 1) return map[digits[0]]

    return recursive(map, digits)
};

function recursive(map, digits, index = 0, result = []) {
    if (!digits[index]) return result
    let newResult = []
    // first time into recursive
    if (!result.length) {
        newResult = [...map[digits[index]]]
    }
    else {
        // combine each letter of number with result generated before
        map[digits[index]].forEach(char => {
            result.forEach(chars => {
                newResult.push(chars + char)
            })
        })
    }
    return recursive(map, digits, index + 1, newResult)
}


// dfs
var letterCombinations = function (digits) {
    if (digits.length === 0) return []
    const res = []
    const map = {
        2: ['a', 'b', 'c'],
        3: ['d', 'e', 'f'],
        4: ['g', 'h', 'i'],
        5: ['j', 'k', 'l'],
        6: ['m', 'n', 'o'],
        7: ['p', 'q', 'r', 's'],
        8: ['t', 'u', 'v'],
        9: ['w', 'x', 'y', 'z'],
    }

    if (digits.length === 1) return map[digits[0]]
    function _dfs(str, index) {
        if(str.length === digits.length) {
            res.push(str)
            return
        }

        for(let char of map[digits[index]]) {
            _dfs(str + char, ++index)
        }
    }

    _dfs("", 0)
    
};
