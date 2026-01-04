package arraysAndHashing

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

func groupAnagrams2(strs []string) [][]string {
	// 遍歷每一個 string
	// 記錄每一個 string 的每個字母出現次數在另一個 hashmap, key: word count arr, value: arr
	// 如果下一個 string 的出現次數跟前面記錄的相同, 則 append 進去
	// 最後把 hashmap 轉成 2d array
	ret := make([][]string, 0)
	m := make(map[[26]int][]string)
	for _, str := range strs {
		count := [26]int{}
		for _, c := range str { // for range string 的每一個 value 是 utf-8 的碼點
			count[c-'a']++
		}

		if m[count] == nil {
			m[count] = make([]string, 0)
		}

		m[count] = append(m[count], str)
	}

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}
