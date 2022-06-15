/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var backspaceCompare = function(s, t) {
    const _backspace = str => str.substring(0, str.length - 1)
    const resFromS = s.split('').reduce((res, char) => {
        if(char === '#') return _backspace(res)
        return res + char
    }, '')
    const resFromT = t.split('').reduce((res, char) => {
        if(char === '#') return _backspace(res)
        return res + char
    }, '')
    
    return resFromS === resFromT
};