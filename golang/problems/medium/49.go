package medium

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	m := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		s := SelectionSort(strs[i])
		m[s] = append(m[s], strs[i])
	}

	ret := make([][]string, 0, len(strs))
	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

func SelectionSort(str string) string {
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

func InsertionSort(str string) string {
	s := []byte(str)
	for i := 0; i < len(s); i++ {
		insertI := s[i]
		idx := i - 1
		for j := i - 1; j >= 0; j-- {
			if insertI < s[j] {
				s[j+1] = s[j]
				idx = idx - 1
			}
		}

		s[idx+1] = insertI
	}

	return string(s)
}
