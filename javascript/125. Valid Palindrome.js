var isPalindrome = function(s) {
    const isNumber = s => s.charCodeAt() >= 48 && s.charCodeAt() <= 57
    const isLetter = s => s.charCodeAt() >= 97 && s.charCodeAt() <= 122
    let lowercase = s.toLowerCase()
    
    function _removeNonAlphanumeric(s) {
        return s.slice().split('').filter(word => isNumber(word) || isLetter(word)).join('')
    }
    
    let onlyAlphaOrNum = _removeNonAlphanumeric(lowercase)
    if(onlyAlphaOrNum.length === 0 || onlyAlphaOrNum.length === 1) return true
    let left = 0, right = onlyAlphaOrNum.length - 1
    while(left <= right) {
        if(onlyAlphaOrNum[left] !== onlyAlphaOrNum[right]) return false;
        left ++;
        right --;
    }
    
    return true;
};