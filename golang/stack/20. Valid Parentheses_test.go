package stack

import (
	"log"
	"testing"
)

func TestIsValid(t *testing.T) {
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

func TestIsValid2(t *testing.T) {
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

func TestIsValidTwoPointer(t *testing.T) {
	case1 := "()"
	ans1 := isValidTwoPointerFailed(case1)
	log.Printf("ans1: %v", ans1)
	case2 := "()[]{}"
	ans2 := isValidTwoPointerFailed(case2)
	log.Printf("ans2: %v", ans2)
	case3 := "(]"
	ans3 := isValidTwoPointerFailed(case3)
	log.Printf("ans3: %v", ans3)
	case4 := "([])"
	ans4 := isValidTwoPointerFailed(case4)
	log.Printf("ans4: %v", ans4)
	case5 := "([)]"
	ans5 := isValidTwoPointerFailed(case5)
	log.Printf("ans5: %v", ans5)
}
