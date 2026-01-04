package arraysAndHashing

import (
	"log"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	ans1 := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	log.Println("ans1: ", ans1)
	ans2 := groupAnagrams([]string{""})
	log.Println("ans2: ", ans2)
	ans3 := groupAnagrams([]string{"a"})
	log.Println("ans3: ", ans3)
}

func TestGroupAnagrams2(t *testing.T) {
	ans1 := groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	log.Println("ans1: ", ans1)
	ans2 := groupAnagrams2([]string{""})
	log.Println("ans2: ", ans2)
	ans3 := groupAnagrams2([]string{"a"})
	log.Println("ans3: ", ans3)
}
