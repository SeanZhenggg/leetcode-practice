package arraysAndHashing

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

type groupAnagramsCase struct {
	input    []string
	expected [][]string
}

var groupAnagramsCases = []groupAnagramsCase{
	{
		input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		expected: [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
	},
	{
		input:    []string{""},
		expected: [][]string{{""}},
	},
	{
		input:    []string{"a"},
		expected: [][]string{{"a"}},
	},
}

func normalizeGroups(groups [][]string) [][]string {
	result := make([][]string, len(groups))
	for i, g := range groups {
		sorted := make([]string, len(g))
		copy(sorted, g)
		sort.Strings(sorted)
		result[i] = sorted
	}
	sort.Slice(result, func(i, j int) bool {
		return strings.Join(result[i], ",") < strings.Join(result[j], ",")
	})
	return result
}

func TestGroupAnagrams(t *testing.T) {
	for _, c := range groupAnagramsCases {
		ans := normalizeGroups(groupAnagrams(c.input))
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestGroupAnagrams2(t *testing.T) {
	for _, c := range groupAnagramsCases {
		ans := normalizeGroups(groupAnagrams2(c.input))
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
