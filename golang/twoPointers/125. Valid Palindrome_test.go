package twoPointers

import (
	"log"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	ans1 := isPalindrome("A man, a plan, a canal: Panama")
	log.Println("ans1: ", ans1)
	ans2 := isPalindrome("race a car")
	log.Println("ans2: ", ans2)
	ans3 := isPalindrome(" ")
	log.Println("ans3: ", ans3)
}
