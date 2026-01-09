package arraysAndHashing

import (
	"log"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	ans1 := productExceptSelf([]int{1, 2, 3, 4})
	log.Println("ans1: ", ans1)
	ans2 := productExceptSelf([]int{-1, 1, 0, -3, 3})
	log.Println("ans2: ", ans2)
}

func TestProductExceptSelf2(t *testing.T) {
	ans3 := productExceptSelf2([]int{1, 2, 3, 4})
	log.Println("ans3: ", ans3)
	ans4 := productExceptSelf2([]int{-1, 1, 0, -3, 3})
	log.Println("ans4: ", ans4)
}

func TestProductExceptSelf3(t *testing.T) {
	ans5 := productExceptSelf3([]int{1, 2, 3, 4})
	log.Println("ans5: ", ans5)
	ans6 := productExceptSelf3([]int{-1, 1, 0, -3, 3})
	log.Println("ans6: ", ans6)
}

func TestProductExceptSelf4(t *testing.T) {
	ans5 := productExceptSelf4([]int{1, 2, 3, 4})
	log.Println("ans5: ", ans5)
	ans6 := productExceptSelf4([]int{-1, 1, 0, -3, 3})
	log.Println("ans6: ", ans6)
}

func TestProductExceptSelf5(t *testing.T) {
	ans5 := productExceptSelf5([]int{1, 2, 3, 4})
	log.Println("ans5: ", ans5)
	ans6 := productExceptSelf5([]int{-1, 1, 0, -3, 3})
	log.Println("ans6: ", ans6)
}

func TestProductExceptSelfReview(t *testing.T) {
	ans1 := productExceptSelfReview([]int{1, 2, 3, 4})
	log.Println("ans1: ", ans1)
	ans2 := productExceptSelfReview([]int{-1, 1, 0, -3, 3})
	log.Println("ans2: ", ans2)
	ans3 := productExceptSelfReview([]int{8, 2, -1, 3, 5})
	log.Println("ans3: ", ans3)
}

func TestProductExceptSelfReview2(t *testing.T) {
	ans1 := productExceptSelfReview2([]int{1, 2, 3, 4})
	log.Println("ans1: ", ans1)
	ans2 := productExceptSelfReview2([]int{-1, 1, 0, -3, 3})
	log.Println("ans2: ", ans2)
	ans3 := productExceptSelfReview2([]int{8, 2, -1, 3, 5})
	log.Println("ans3: ", ans3)
}
