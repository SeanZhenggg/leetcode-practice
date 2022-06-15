/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */

// wrong answer
class Stack {
    constructor() {
        this.stack = []
        return this
    }

    push(x) {
        this.stack.push(x)
    }

    pop() {
        if (!this.isEmpty()) {
            return this.stack.pop()
        }
        return null
    }

    top() {
        const last = this.stack.pop()
        this.stack.push(last)
        return last
    }

    isEmpty() {
        return this.stack.length === 0
    }

    getStack() {
        return this.stack
    }
}

var removeDuplicates = function (s, k) {
    if (s.length < k) return s

    const stack = new Stack()
    const valMap = new Map()
    const countMap = new Map()
    for (let i = 0; i < s.length; i++) {
        if (i === 0 || stack.top() !== s[i]) {
            countMap.set(s[i], 1)
        }
        else {
            let countVal = countMap.get(s[i])
            countMap.set(s[i], ++countVal)
        }

        if (!valMap.has(s[i])) {
            valMap.set(s[i], 1)
        }
        else {
            let val = valMap.get(s[i])
            valMap.set(s[i], ++val)
        }

        stack.push(s[i])

        if (countMap.get(s[i]) === k && valMap.get(s[i]) >= k) {
            for (let i = 1; i <= k; i++) stack.pop()
            const val = valMap.get(s[i])
            valMap.set(s[i], val - k)
            countMap.set(s[i], val - k)
            count = countMap.get(stack.top())
        }
    }
    return stack.getStack().join('')
};

// solution 1
var removeDuplicates = function(s, k) {
    var stack = [], result = ''
    
    for(let i = 0; i< s.length; i++) {
        if(stack.length > 0 && stack[stack.length - 1]['char'] === s[i]) {
            stack[stack.length - 1]['count'] += 1;
            if(stack[stack.length - 1]['count'] >= k) {
                stack.pop()
            }
        } else {
            stack.push({ char: s[i], count: 1 })   
        }
    }
    
    for(let j = 0; j < stack.length; j++) {
        result += stack[j]['char'].repeat(stack[j]['count'])
    }
    
    return result;
};

// solution 2
var removeDuplicates = function(S, K) {
    let SC = S.split(""), st = [0], i, j
    for (i = 1, j = 1; j < S.length; SC[++i] = SC[++j]) {
        if (SC[i] !== SC[i-1]) st.push(i)
        else if (i - st[st.length-1] + 1 === K) i = st.pop()-1
    }
    return SC.slice(0,i+1).join("")
};