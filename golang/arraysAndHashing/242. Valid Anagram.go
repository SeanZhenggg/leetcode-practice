package arraysAndHashing

import "log"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func TestIsAnagram() {
	ans1 := isAnagram("anagram", "nagaram")
	log.Println("ans1: ", ans1)
	ans2 := isAnagram("rat", "car")
	log.Println("ans2: ", ans2)
}
