package stack

import (
	"testing"
)

type EvalRPNCase struct {
	input    []string
	expected int
}

var evalRPNCaseCases = []EvalRPNCase{
	{input: []string{"2", "1", "+", "3", "*"}, expected: 9},
	{input: []string{"4", "13", "5", "/", "+"}, expected: 6},
	{input: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, expected: 22},
}

func TestEvalRPN(t *testing.T) {
	for _, c := range evalRPNCaseCases {
		ans := evalRPN(c.input)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestEvalRPNBruteForce(t *testing.T) {
	for _, c := range evalRPNCaseCases {
		ans := evalRPNBruteForce(c.input)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
