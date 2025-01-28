package arraysAndHashing

import "log"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	m := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		s := selectionSort(strs[i])
		m[s] = append(m[s], strs[i])
	}

	ret := make([][]string, 0, len(strs))
	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

func selectionSort(str string) string {
	s := []byte(str)
	for i := 0; i < len(s); i++ {
		minI := i
		for j := i; j < len(s); j++ {
			if s[j] < s[minI] {
				minI = j
			}
		}

		s[minI], s[i] = s[i], s[minI]
	}

	return string(s)
}

func Test_GroupAnagrams() {
	ans1 := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	log.Println("ans1: ", ans1)
	ans2 := groupAnagrams([]string{""})
	log.Println("ans2: ", ans2)
	ans3 := groupAnagrams([]string{"a"})
	log.Println("ans3: ", ans3)
}
