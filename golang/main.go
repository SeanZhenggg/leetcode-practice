package main

import (
	"golang/problems/medium"
	"log"
)

func main() {
	// 347 test
	//ans1 := medium.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	//log.Println("ans1: ", ans1)
	//ans2 := medium.TopKFrequent([]int{1}, 1)
	//log.Println("ans2: ", ans2)

	// 242 test
	//ans1 := easy.IsAnagram("anagram", "nagaram")
	//log.Println("ans1: ", ans1)
	//ans2 := easy.IsAnagram("rat", "car")
	//log.Println("ans2: ", ans2)

	// 1 test
	//ans1 := easy.TwoSum([]int{2, 7, 11, 15}, 9)
	//log.Println("ans1: ", ans1)
	//ans2 := easy.TwoSum([]int{3, 2, 4}, 6)
	//log.Println("ans2: ", ans2)
	//ans3 := easy.TwoSum([]int{3, 3}, 6)
	//log.Println("ans3: ", ans3)

	// 49 test
	ans1 := medium.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	log.Println("ans1: ", ans1)
	ans2 := medium.GroupAnagrams([]string{""})
	log.Println("ans2: ", ans2)
	ans3 := medium.GroupAnagrams([]string{"a"})
	log.Println("ans3: ", ans3)
}
