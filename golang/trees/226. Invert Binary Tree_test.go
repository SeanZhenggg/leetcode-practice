package trees

import "testing"

type invertTreeCase struct {
	input    []string
	expected []string
}

var invertTreeCases = []invertTreeCase{
	{
		input:    []string{"4", "2", "7", "1", "3", "6", "9"},
		expected: []string{"4", "7", "2", "9", "6", "3", "1"},
	},
	{
		input:    []string{"2", "1", "3"},
		expected: []string{"2", "3", "1"},
	},
	{
		input:    []string{},
		expected: []string{},
	},
}

func Test_InvertTree(t *testing.T) {
	for _, c := range invertTreeCases {
		result := invertTree(generateTree(c.input))
		expected := generateTree(c.expected)
		if !isSameTree(result, expected) {
			t.Errorf("inverted tree does not match expected")
		}
	}
}

func Test_InvertTree2(t *testing.T) {
	for _, c := range invertTreeCases {
		result := invertTree2(generateTree(c.input))
		expected := generateTree(c.expected)
		if !isSameTree(result, expected) {
			t.Errorf("inverted tree does not match expected")
		}
	}
}

func Test_InvertTreeReview1(t *testing.T) {
	for _, c := range invertTreeCases {
		result := invertTreeReview1(generateTree(c.input))
		expected := generateTree(c.expected)
		if !isSameTree(result, expected) {
			t.Errorf("inverted tree does not match expected")
		}
	}
}

func Test_InvertTreeReview2(t *testing.T) {
	for _, c := range invertTreeCases {
		result := invertTreeReview2(generateTree(c.input))
		expected := generateTree(c.expected)
		if !isSameTree(result, expected) {
			t.Errorf("inverted tree does not match expected")
		}
	}
}

func Test_InvertTreeReview3(t *testing.T) {
	for _, c := range invertTreeCases {
		result := invertTreeReview3(generateTree(c.input))
		expected := generateTree(c.expected)
		if !isSameTree(result, expected) {
			t.Errorf("inverted tree does not match expected")
		}
	}
}

func Test_InvertTreeReview4(t *testing.T) {
	for _, c := range invertTreeCases {
		result := invertTreeReview4(generateTree(c.input))
		expected := generateTree(c.expected)
		if !isSameTree(result, expected) {
			t.Errorf("inverted tree does not match expected")
		}
	}
}
