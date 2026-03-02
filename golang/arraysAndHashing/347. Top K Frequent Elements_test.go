package arraysAndHashing

import (
	"reflect"
	"sort"
	"testing"
)

type topKFrequentCase struct {
	nums     []int
	k        int
	expected []int
}

var topKFrequentCases = []topKFrequentCase{
	{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expected: []int{1, 2}},
	{nums: []int{1}, k: 1, expected: []int{1}},
	{nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, k: 2, expected: []int{1, 2}},
	{nums: []int{-5, -2, -1, 2, 5, -5, 2, -1}, k: 3, expected: []int{-5, -1, 2}},
}

func TestTopKFrequent(t *testing.T) {
	for _, c := range topKFrequentCases {
		ans := topKFrequent(c.nums, c.k)
		sort.Ints(ans)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestTopKFrequent2(t *testing.T) {
	for _, c := range topKFrequentCases {
		ans := topKFrequent2(c.nums, c.k)
		sort.Ints(ans)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestTopKFrequent3(t *testing.T) {
	for _, c := range topKFrequentCases {
		ans := topKFrequent3(c.nums, c.k)
		sort.Ints(ans)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestTopKFrequent4(t *testing.T) {
	for _, c := range topKFrequentCases {
		ans := topKFrequent4(c.nums, c.k)
		sort.Ints(ans)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
