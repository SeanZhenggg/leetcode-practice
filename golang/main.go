package main

import (
	"golang/problems/easy"
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
	//ans1 := medium.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	//log.Println("ans1: ", ans1)
	//ans2 := medium.GroupAnagrams([]string{""})
	//log.Println("ans2: ", ans2)
	//ans3 := medium.GroupAnagrams([]string{"a"})
	//log.Println("ans3: ", ans3)

	// 238 test
	//ans1 := medium.ProductExceptSelf([]int{1, 2, 3, 4})
	//log.Println("ans1: ", ans1)
	//ans2 := medium.ProductExceptSelf([]int{-1, 1, 0, -3, 3})
	//log.Println("ans2: ", ans2)
	//log.Println()
	//log.Println()
	//ans3 := medium.ProductExceptSelf2([]int{1, 2, 3, 4})
	//log.Println("ans3: ", ans3)
	//ans4 := medium.ProductExceptSelf2([]int{-1, 1, 0, -3, 3})
	//log.Println("ans4: ", ans4)
	//log.Println()
	//log.Println()
	//ans5 := medium.ProductExceptSelf3([]int{1, 2, 3, 4})
	//log.Println("ans5: ", ans5)
	//ans6 := medium.ProductExceptSelf3([]int{-1, 1, 0, -3, 3})
	//log.Println("ans6: ", ans6)

	// 125 test
	ans1 := easy.IsPalindrome("A man, a plan, a canal: Panama")
	log.Println("ans1: ", ans1)
	ans2 := easy.IsPalindrome("race a car")
	log.Println("ans2: ", ans2)
	ans3 := easy.IsPalindrome(" ")
	log.Println("ans3: ", ans3)
}
