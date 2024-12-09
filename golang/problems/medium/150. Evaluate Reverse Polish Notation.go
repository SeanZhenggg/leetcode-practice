package medium

import (
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	operators := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	st := make([]int, 0, len(tokens))

	for _, v := range tokens {
		var val int
		if _, found := operators[v]; found {
			if len(st) > 0 {
				op1 := st[len(st)-1]
				st = st[:len(st)-1]
				op2 := st[len(st)-1]
				st = st[:len(st)-1]
				switch v {
				case "+":
					val = op2 + op1
				case "-":
					val = op2 - op1
				case "*":
					val = op2 * op1
				case "/":
					val = op2 / op1
				}
			}
		} else {
			val, _ = strconv.Atoi(v)
		}

		st = append(st, val)
	}

	return st[0]
}

func Test_EvalRPN() {
	case1 := []string{"2", "1", "+", "3", "*"}
	ans1 := evalRPN(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []string{"4", "13", "5", "/", "+"}
	ans2 := evalRPN(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	ans3 := evalRPN(case3)
	log.Printf("ans3: %v", ans3)

}
