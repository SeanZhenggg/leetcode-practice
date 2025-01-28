package stack

import "log"

func isValid(s string) bool {
	st := make([]string, 0, len(s))
	parenthesesMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, ch := range s {
		sc := string(ch)
		if len(st) > 0 && parenthesesMap[sc] == st[len(st)-1] {
			st = st[:len(st)-1]
		} else {
			st = append(st, sc)
		}
	}

	return len(st) == 0
}

func isValid2(s string) bool {
	st := make([]rune, 0, len(s))
	parenthesesMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if open, found := parenthesesMap[ch]; found {
			if len(st) > 0 && open == st[len(st)-1] {
				st = st[:len(st)-1]
			} else {
				return false
			}
		} else {
			st = append(st, ch)
		}
	}

	return len(st) == 0
}

func Test_IsValid() {
	case1 := "()"
	ans1 := isValid(case1)
	log.Printf("ans1: %v", ans1)
	case2 := "()[]{}"
	ans2 := isValid(case2)
	log.Printf("ans2: %v", ans2)
	case3 := "(]"
	ans3 := isValid(case3)
	log.Printf("ans3: %v", ans3)
	case4 := "([])"
	ans4 := isValid(case4)
	log.Printf("ans4: %v", ans4)
	case5 := "([)]"
	ans5 := isValid(case5)
	log.Printf("ans5: %v", ans5)
}

func Test_IsValid2() {
	case1 := "()"
	ans1 := isValid2(case1)
	log.Printf("ans1: %v", ans1)
	case2 := "()[]{}"
	ans2 := isValid2(case2)
	log.Printf("ans2: %v", ans2)
	case3 := "(]"
	ans3 := isValid2(case3)
	log.Printf("ans3: %v", ans3)
	case4 := "([])"
	ans4 := isValid2(case4)
	log.Printf("ans4: %v", ans4)
	case5 := "([)]"
	ans5 := isValid2(case5)
	log.Printf("ans5: %v", ans5)
}
