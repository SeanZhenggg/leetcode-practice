package arraysAndHashing

import (
	"log"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	ans1 := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	log.Println("ans1: ", ans1)
	ans2 := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	log.Println("ans2: ", ans2)
	ans3 := longestConsecutive([]int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161})
	log.Println("ans3: ", ans3)
	ans4 := longestConsecutive2([]int{1, 0, 1, 2})
	log.Println("ans4: ", ans4)
}

func TestLongestConsecutive2(t *testing.T) {
	ans1 := longestConsecutive2([]int{100, 4, 200, 1, 3, 2})
	log.Println("ans1: ", ans1)
	ans2 := longestConsecutive2([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	log.Println("ans2: ", ans2)
	ans3 := longestConsecutive2([]int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161})
	log.Println("ans3: ", ans3)
	ans4 := longestConsecutive2([]int{1, 0, 1, 2})
	log.Println("ans4: ", ans4)
}

func TestLongestConsecutive3(t *testing.T) {
	ans1 := longestConsecutive3([]int{100, 4, 200, 1, 3, 2})
	log.Println("ans1: ", ans1)
	ans2 := longestConsecutive3([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	log.Println("ans2: ", ans2)
	ans3 := longestConsecutive3([]int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161})
	log.Println("ans3: ", ans3)
	ans4 := longestConsecutive3([]int{3, 4, 5, 4})
	log.Println("ans4: ", ans4)
}

func TestLongestConsecutive4(t *testing.T) {
	ans1 := longestConsecutive4([]int{100, 4, 200, 1, 3, 2})
	log.Println("ans1: ", ans1)
	ans2 := longestConsecutive4([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	log.Println("ans2: ", ans2)
	ans3 := longestConsecutive4([]int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161})
	log.Println("ans3: ", ans3)
	ans4 := longestConsecutive4([]int{3, 4, 5, 4})
	log.Println("ans4: ", ans4)
}
